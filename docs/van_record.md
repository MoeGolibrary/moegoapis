# Van Staff Record API 文档（`moego.business.van_record.v1`）

## 1. 功能概述

`Van Staff Record` 模块提供按天查询 van 与 staff 历史绑定快照的能力。

该接口用于：

- 查询某一天某些 van 上绑定了哪些 staff。
- 保留 van 和 staff 的原始产品 ID，方便做下游映射。
- 支撑按 van 维度做历史归因、排查和报表分析。

---

## 2. 设计目标

- **历史快照**：返回指定日期的 van 与 staff 绑定状态。
- **Van 维度查询**：请求按 `van_ids` 收敛，而不是按 business 全量扫描。
- **对外分页规范**：接口仍然遵循 OpenAPI list 类接口的分页规范。
- **日期可读性**：使用 `yyyy-MM-dd`，不暴露底层时间戳实现细节。
- **桥接友好**：返回 `rawId`，便于外部系统继续关联下游数据。

---

## 3. 核心模型

### 3.1 VanStaffRecord

表示某一天某一辆 van 的 staff 绑定快照。

| 字段名 | 类型 | 说明 |
|--------|------|------|
| `vanId` | string | van 的混淆 ID |
| `vanRawId` | int64 | van 在 MoeGo 产品中的原始数值 ID |
| `businessId` | string | van 所属 business 的混淆 ID |
| `date` | string | 快照日期，格式为 `yyyy-MM-dd` |
| `assignedStaffs` | Array(`AssignedStaff`) | 该日期下绑定到该 van 的 staff 列表 |

### 3.2 AssignedStaff

复用 `moego.business.van.v1` 中的 `AssignedStaff` 模型。

| 字段名 | 类型 | 说明 |
|--------|------|------|
| `id` | string | staff 的混淆 ID |
| `rawId` | int64 | staff 在 MoeGo 产品中的原始数值 ID |
| `name` | string | staff 展示名称 |

---

## 4. 典型使用场景

### 场景：按 van 归因某天的业务数据

1. 调用 `ListVanStaffRecords`，传入目标 `van_ids` 和目标日期。
2. 对每条返回记录，识别该 van 在当天绑定的 staff。
3. 使用这些 staff ID 去查询 appointment、service、revenue 等下游数据。
4. 最终按 `vanId` 或 `vanRawId` 聚合结果。

---

## 5. API 接口说明

### 5.1 查询 Van Staff 历史记录（`ListVanStaffRecords`）

- **方法名**：`ListVanStaffRecords`
- **HTTP Method**：POST
- **Path**：`/v1/van_staff_records:list`

#### 功能说明

按请求中的 `van_ids` 和指定日期，返回这些 van 在该日期下的历史 staff 绑定快照。

#### 使用场景

- 分析某些 van 在某一天的 staff 贡献情况。
- 构建 van 维度的历史收入归因链路。
- 排查历史 staff 绑定变化。

#### 请求参数

| 字段名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `pagination` | Pagination | 是 | 分页参数 |
| `vanIds` | Array(string) | 是 | 目标 van 的混淆 ID 列表 |
| `date` | string | 是 | 快照日期，格式 `yyyy-MM-dd` |

#### 返回值

| 字段名 | 类型 | 说明 |
|--------|------|------|
| `nextPageToken` | string | 下一页分页 token |
| `records` | Array(`VanStaffRecord`) | 命中的历史快照记录 |

#### 返回行为

- 返回结果在分页后按 `van_id` 升序排列。

#### 错误码

| 错误码 | 说明 |
|--------|------|
| `INVALID_ARGUMENT` | `van_ids`、`date` 或 `pagination` 不合法 |
| `PERMISSION_DENIED` | 调用方无权访问请求中的一个或多个 van |

---

## 6. 使用示例

### 示例：查询某一天多个 van 的历史绑定记录

**请求**

```http
POST /v1/van_staff_records:list
Content-Type: application/json

{
  "pagination": {
    "pageSize": 50,
    "pageToken": "1"
  },
  "vanIds": ["van_001", "van_002"],
  "date": "2026-06-17"
}
```

**响应**

```json
{
  "nextPageToken": "",
  "records": [
    {
      "vanId": "van_001",
      "vanRawId": 101,
      "businessId": "biz_001",
      "date": "2026-06-17",
      "assignedStaffs": [
        {
          "id": "stf_001",
          "rawId": 1001,
          "name": "Jane Doe"
        },
        {
          "id": "stf_002",
          "rawId": 1002,
          "name": "John Smith"
        }
      ]
    },
    {
      "vanId": "van_002",
      "vanRawId": 102,
      "businessId": "biz_001",
      "date": "2026-06-17",
      "assignedStaffs": [
        {
          "id": "stf_003",
          "rawId": 1003,
          "name": "Alice Lee"
        }
      ]
    }
  ]
}
```
