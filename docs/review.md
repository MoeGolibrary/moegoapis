# 📝 Review API Documentation (`moego.business.review.v1`)

## 📌 1. Functional Overview

The **Review** module provides APIs to manage customer feedback for services provided at your business.

Each review is associated with a specific appointment and can include ratings and comments about multiple staff members
and pets. This interface enables:

- Retrieving customer reviews based on specified criteria.
- Filtering reviews by feedback source (e.g., SMS, grooming report, portal).
- Filtering reviews by staff members involved in the service.
- Paginating through large datasets of reviews for scalable analysis.

Useful for scenarios such as analyzing customer satisfaction, evaluating staff performance, and improving service
quality.

---

## 🎯 2. Design Goals

- **Centralized Feedback Management**: Provides a unified way to collect and analyze customer feedback.
- **Rich Data Model**: Supports complex relationships like ratings, comments, and associations with appointments, staff,
  and pets.
- **Secure and Reliable**: Ensures access control and data integrity.
- **Easy Integration**: Offers RESTful interfaces compatible with mainstream development languages and frameworks.

Applicable to scenarios like:

- Tracking customer satisfaction trends over time.
- Evaluating individual staff performance based on reviews.
- Analyzing feedback from different sources (SMS, portal, reports).

---

## 🧩 3. Core Concepts

### 1. Review

Represents customer feedback submitted after receiving services.

| Field Name          | Type          | Description                                    |
|---------------------|---------------|------------------------------------------------|
| `id`                | string        | Unique identifier (e.g., `"rev_001"`).         |
| `customer_id`       | string        | ID of the customer who submitted the review    |
| `appointment_id`    | string        | ID of the appointment being reviewed           |
| `staff_ids`         | Array(string) | IDs of staff members included in the review    |
| `pet_ids`           | Array(string) | IDs of pets involved in the reviewed service   |
| `source`            | Source        | Channel through which the review was submitted |
| `score`             | uint32        | Customer satisfaction rating (scale: 1–5)      |
| `content`           | string        | Detailed feedback content                      |
| `review_time`       | Timestamp     | When the review was submitted by the customer  |
| `created_time`      | Timestamp     | When this review was created in the system     |
| `last_updated_time` | Timestamp     | When this review was last modified             |

#### Enum: Source

Describes how the review was submitted by the customer.

- `SOURCE_UNSPECIFIED`
- `SMS`
- `GROOMING_REPORT`
- `PET_PARENT_PORTAL`

---

## 📈 4. Typical Usage Flow

### ✅ Scenario: User Integrates and Debugs Review API

Here is a typical integration flow:

1. **List Reviews**
    - Retrieve all reviews for a specific business.
    - Optionally filter by source or staff member.

2. **Monitoring & Analysis**
    - Regularly fetch reviews to monitor customer satisfaction.
    - Filter by staff to evaluate performance.
    - Analyze feedback trends by source (e.g., SMS vs. Portal).

3. **Integration with Dashboards**
    - Display aggregated review scores.
    - Highlight top/bottom performing staff based on feedback.

---

## 🛠️ 5. API Interface Descriptions

### 1. List Reviews (`ListReviews`)

- **Method**: `ListReviews`
- **HTTP Method**: POST
- **Path**: `/v1/reviews:list`

#### ✅ Functionality:

Retrieves a paginated list of customer reviews based on specified criteria.

Supports filtering by feedback source and staff members to facilitate targeted analysis of customer satisfaction and
service quality.

#### 🎯 Use Cases:

- View all reviews for a business location.
- Filter reviews by source (e.g., only those from SMS).
- Filter reviews related to specific staff members.

#### 🔧 Request Parameters:

| Field Name         | Type          | Required | Description                             |
|--------------------|---------------|----------|-----------------------------------------|
| `pagination`       | Pagination    | Yes      | Page size and token                     |
| `business_id`      | string        | Yes      | Business location ID to scope reviews   |
| `filter.sources`   | Array(Source) | No       | Feedback channels to include in results |
| `filter.staff_ids` | Array(string) | No       | Staff members to filter reviews by      |

#### 📌 Return Value:

| Field Name        | Type          | Description                                   |
|-------------------|---------------|-----------------------------------------------|
| `next_page_token` | string        | Token for retrieving the next page of results |
| `reviews`         | Array(Review) | List of reviews matching the request criteria |

#### ⚠️ Error Codes:

- `INVALID_ARGUMENT`: Missing or invalid pagination parameters.
- `PERMISSION_DENIED`: Caller lacks access rights.
- `NOT_FOUND`: The business ID doesn't exist.

---

## 🧪 6. Usage Examples

### Example 1: List All Reviews for a Business

**Request**

```http
POST /v1/reviews:list
Content-Type: application/json

{
  "pagination": {
    "page_size": 20
  },
  "business_id": "bus_001"
}
```

**Response**

```json
{
  "next_page_token": "CBAQAA==",
  "reviews": [
    {
      "id": "rev_001",
      "customer_id": "cus_001",
      "appointment_id": "apt_001",
      "staff_ids": [
        "stf_001"
      ],
      "pet_ids": [
        "pet_001"
      ],
      "source": "PET_PARENT_PORTAL",
      "score": 5,
      "content": "Excellent service! My dog looked amazing.",
      "review_time": "2023-09-15T10:00:00Z",
      "created_time": "2023-09-15T10:05:00Z",
      "last_updated_time": "2023-09-15T10:05:00Z"
    },
    {
      "id": "rev_002",
      "customer_id": "cus_002",
      "appointment_id": "apt_002",
      "staff_ids": [
        "stf_002"
      ],
      "pet_ids": [
        "pet_002"
      ],
      "source": "SMS",
      "score": 4,
      "content": "Good experience overall.",
      "review_time": "2023-09-16T11:00:00Z",
      "created_time": "2023-09-16T11:05:00Z",
      "last_updated_time": "2023-09-16T11:05:00Z"
    }
  ]
}
```

---

### Example 2: Filter Reviews by Source and Staff

**Request**

```http
POST /v1/reviews:list
Content-Type: application/json

{
  "pagination": {
    "page_size": 20
  },
  "business_id": "bus_001",
  "filter": {
    "sources": ["SMS", "PET_PARENT_PORTAL"],
    "staff_ids": ["stf_001"]
  }
}
```

**Response**

```json
{
  "next_page_token": "",
  "reviews": [
    {
      "id": "rev_001",
      "customer_id": "cus_001",
      "appointment_id": "apt_001",
      "staff_ids": [
        "stf_001"
      ],
      "pet_ids": [
        "pet_001"
      ],
      "source": "PET_PARENT_PORTAL",
      "score": 5,
      "content": "Excellent service! My dog looked amazing.",
      "review_time": "2023-09-15T10:00:00Z",
      "created_time": "2023-09-15T10:05:00Z",
      "last_updated_time": "2023-09-15T10:05:00Z"
    }
  ]
}
```

---

## ⚠️ 7. Usage Limitations

TODO

---

## ❓ 8. FAQ

| Question                                             | Answer                                                                               |
|------------------------------------------------------|--------------------------------------------------------------------------------------|
| How do I retrieve all reviews for a business?        | Use `ListReviews` with the `business_id` parameter.                                  |
| Can I filter reviews by multiple staff members?      | Yes, provide an array of staff IDs in the `filter.staff_ids` field.                  |
| How are reviews sorted?                              | Results are typically sorted by `review_time` descending unless otherwise specified. |
| Why does listing reviews return "permission denied"? | Ensure you have the correct access rights for this operation.                        |
| How do I handle large result sets efficiently?       | Use pagination via `page_size` and `page_token`.                                     |

---

## 📌 9. Common Error Codes

| Error Code          | Description                          |
|---------------------|--------------------------------------|
| `NOT_FOUND`         | The business ID doesn't exist.       |
| `PERMISSION_DENIED` | Caller lacks access rights.          |
| `INVALID_ARGUMENT`  | Invalid request parameters provided. |
| `INTERNAL`          | Internal server error occurred.      |

---

## 📎 10. Related File References

- [review.proto](../moego/business/review/v1/review.proto)
- [review_service.proto](../moego/business/review/v1/review_service.proto)
- [pagination.proto](../moego/common/v1/pagination.proto)