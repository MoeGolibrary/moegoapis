"""
Moego Orders Export Script
==========================

This script exports order data from the Moego API to a CSV file.

Features:
- Automatically retrieves company and business IDs
- Fetches orders within a specified time range
- Exports data to CSV format
- Handles pagination for large datasets
- Includes security checks for URLs and filenames

Usage:
python list_order.py --api_key YOUR_API_KEY [--start_time START_TIME] [--end_time END_TIME]

Time format should be ISO format: YYYY-MM-DDTHH:MM:SSZ
If no time range is specified, it defaults to the last 7 days.
"""

import requests
import csv
import time
import argparse
from datetime import datetime, timedelta
from urllib.parse import urlparse

# === CONFIG ===

# === STEP 1: GET COMPANY & BUSINESS IDs ===
def get_company_and_business_ids(headers):
    """
    Call companies:list and businesses:list to get the correct obfuscated IDs
    """
    print("Fetching company and business IDs...")

    # --- Step 1: Get Companies ---
    company_url = "https://openapi.moego.pet/v1/companies:list"
    company_body = {
        "pagination": {
            "pageSize": 100,
            "pageToken": "1"
        }
    }

    try:
        # Validate URL safety
        if not is_safe_url(company_url):
            raise ValueError("Unsafe URL detected")
            
        response = requests.post(company_url, headers=headers, json=company_body, timeout=30)
        if response.status_code != 200:
            print("Failed to fetch companies:", response.status_code, response.text)
            return None, []

        data = response.json()
        companies = data.get("companies", [])
        if not companies:
            print("No companies found.")
            return None, []

        # Take the first company (or choose as needed)
        company = companies[0]
        company_id = company["id"]
        print(f"✅ Using companyId: {company_id}")

        # Add tip on how to manually send request to get company IDs
        print(f"\n💡 To manually get company IDs, use the following curl command:")
        print(f"curl -X POST '{company_url}' \\")
        print(f"  -H 'Authorization: Basic YOUR_API_KEY' \\")
        print(f"  -H 'Content-Type: application/json' \\")
        print(f"  -d '{{\"pagination\": {{\"pageSize\": 100, \"pageToken\": \"1\"}}}}'")
        print()

        # --- Step 2: Get Businesses ---
        business_url = "https://openapi.moego.pet/v1/businesses:list"
        business_body = {
            "companyId": company_id,
            "pagination": {
                "pageSize": 100,
                "pageToken": "1"
            }
        }

        # Validate URL safety
        if not is_safe_url(business_url):
            raise ValueError("Unsafe URL detected")
            
        response = requests.post(business_url, headers=headers, json=business_body, timeout=30)
        if response.status_code != 200:
            print("Failed to fetch businesses:", response.status_code, response.text)
            return company_id, []

        data = response.json()
        businesses = data.get("businesses", [])
        business_ids = [biz["id"] for biz in businesses]
        print(f"✅ Found {len(business_ids)} business(es): {business_ids}")

        # Add tip on how to manually send request to get business IDs
        print(f"\n💡 To manually get business IDs, use the following curl command:")
        print(f"curl -X POST '{business_url}' \\")
        print(f"  -H 'Authorization: Basic YOUR_API_KEY' \\")
        print(f"  -H 'Content-Type: application/json' \\")
        print(f"  -d '{{\"companyId\": \"{company_id}\", \"pagination\": {{\"pageSize\": 100, \"pageToken\": \"1\"}}}}'")
        print()

        return company_id, business_ids

    except Exception as e:
        print("Error during ID fetching:", e)
        return None, []


# === FUNCTION TO GET ORDERS ===
def get_all_orders(company_id, business_ids, headers, START_TIME, END_TIME):
    url = "https://openapi.moego.pet/v1/orders:list"
    orders = []
    page_token = "1"

    # Validate URL safety
    if not is_safe_url(url):
        raise ValueError("Unsafe URL detected")
        
    while True:
        body = {
            "companyId": company_id,
            "businessIds": business_ids,
            "pagination": {
                "pageSize": 100,
                "pageToken": page_token
            },
            "filter": {
                "createdTime": {
                    "startTime": START_TIME,
                    "endTime": END_TIME
                }
            }
        }
        if page_token:
            body["pagination"]["pageToken"] = page_token

        response = requests.post(url, headers=headers, json=body)
        print("Status code:", response.status_code)
        print("Response text:", response.text)

        try:
            response = requests.post(url, headers=headers, json=body, timeout=30)
            print("Status code:", response.status_code)
            # Limit response output length to avoid overly long output
            response_text = response.text
            if len(response_text) > 500:
                response_text = response_text[:500] + "...(truncated)"
            print("Response text:", response_text)

            data = response.json()
        except Exception as e:
            print("Failed to parse JSON or request failed:", e)
            break

        if "orders" not in data:
            print("No 'orders' key in response:", data)
            break

        orders.extend(data["orders"])

        if "nextPageToken" in data:
            page_token = data["nextPageToken"]
            time.sleep(1)  # Prevent triggering rate limiting
        else:
            break

        if page_token == "":
            break

    return orders


# === FUNCTION TO WRITE TO CSV ===
def write_orders_to_csv(orders, filename):
    if not orders:
        print("No orders to save.")
        return

    # Validate filename safety
    if not is_safe_filename(filename):
        raise ValueError("Unsafe filename detected")
        
    keys = ["id", "status", "totalAmount", "currency", "createdTime", "updatedTime", "businessId", "clientId"]
    with open(filename, mode="w", newline="", encoding="utf-8") as file:
        writer = csv.DictWriter(file, fieldnames=keys)
        writer.writeheader()
        for order in orders:
            flat_order = {key: order.get(key, "") for key in keys}
            writer.writerow(flat_order)
    print(f"✅ Saved {len(orders)} orders to {filename}")


# === URL Safety Check Function ===
def is_safe_url(url):
    """
    Check if the URL is safe
    """
    try:
        parsed = urlparse(url)
        # Only allow HTTPS protocol and specified domain
        if parsed.scheme != 'https':
            return False
        if parsed.netloc != 'openapi.moego.pet':
            return False
        return True
    except Exception:
        return False


# === Filename Safety Check Function ===
def is_safe_filename(filename):
    """
    Check if the filename is safe
    """
    # Prohibit path traversal characters
    illegal_chars = ['/', '\\', '..', ':']
    for char in illegal_chars:
        if char in filename:
            return False
    return True


# === MAIN RUN ===
if __name__ == "__main__":
    # Create argument parser
    parser = argparse.ArgumentParser(description='Fetch Moego orders data')
    parser.add_argument('--api_key', required=True, help='API Key for authentication')
    # Calculate default time range (last 3 days instead of 7)
    default_end_time = datetime.utcnow().strftime('%Y-%m-%dT%H:%M:%SZ')
    default_start_time = (datetime.utcnow() - timedelta(days=3)).strftime('%Y-%m-%dT%H:%M:%SZ')
    parser.add_argument('--start_time', default=default_start_time, help='Start time in ISO format (default: 3 days ago)')
    parser.add_argument('--end_time', default=default_end_time, help='End time in ISO format (default: now)')
    
    args = parser.parse_args()
    
    # Use argument values
    API_KEY = args.api_key
    START_TIME = args.start_time
    END_TIME = args.end_time
    
    # Validate input parameters
    if not API_KEY:
        print("❌ API Key is required")
        exit(1)
        
    # Validate time format
    try:
        datetime.strptime(START_TIME, '%Y-%m-%dT%H:%M:%SZ')
        datetime.strptime(END_TIME, '%Y-%m-%dT%H:%M:%SZ')
    except ValueError:
        print("❌ Invalid time format. Use ISO format: YYYY-MM-DDTHH:MM:SSZ")
        exit(1)
    
    # Generate CSV filename based on time range
    start_date = START_TIME.replace("T", "_").replace(":", "-").split(".")[0]
    end_date = END_TIME.replace("T", "_").replace(":", "-").split(".")[0]
    OUTPUT_CSV = f"moego_orders_{start_date}_to_{end_date}.csv"
    
    # === HEADERS ===
    headers = {
        "Authorization": f"Basic {API_KEY}",
        "Content-Type": "application/json"
    }

    print("Initializing Moego API data fetch...")

    # Automatically obtain the companyId and Business ID
    company_id, business_ids = get_company_and_business_ids(headers)

    if not company_id or not business_ids:
        print("❌ Failed to retrieve valid IDs. Exiting.")
    else:
        print(f"Fetching orders for companyId={company_id}, businessIds={business_ids}...")
        # Add tip on how to manually send request
        print(f"\n💡 To manually send this request, use the following curl command:")
        # Hide API key to prevent leakage
        print(f"curl -X POST 'https://openapi.moego.pet/v1/orders:list' \\")
        print(f"  -H 'Authorization: Basic YOUR_API_KEY' \\")
        print(f"  -H 'Content-Type: application/json' \\")
        print(f"  -d '{{\"companyId\": \"{company_id}\", \"businessIds\": {business_ids}, \"pagination\": {{\"pageSize\": 100, \"pageToken\": \"1\"}}, \"filter\": {{\"createdTime\": {{\"startTime\": \"{START_TIME}\", \"endTime\": \"{END_TIME}\"}}}}}}'")
        print()
        all_orders = get_all_orders(company_id, business_ids, headers, START_TIME, END_TIME)
        write_orders_to_csv(all_orders, OUTPUT_CSV)