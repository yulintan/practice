# logs = [
#   {"batch_id": 101, "resource_id": "R1", "recorded_at": "2026-05-10T08:00:00Z"},
#   {"batch_id": 101, "resource_id": "R2", "recorded_at": "2026-05-10T08:05:00Z"},
#   {"batch_id": 101, "resource_id": "R1", "recorded_at": "2026-05-10T08:10:00Z"},
#   {"batch_id": 202, "resource_id": "R1", "recorded_at": "2026-05-11T09:00:00Z"},
#   {"batch_id": 202, "resource_id": "R3", "recorded_at": "2026-05-11T09:15:00Z"},
# ]

# Task:
# Write a function that returns a summary per batch with:
# - The number of unique resources utilized
# - The earliest recorded time for that batch

# Expected output shape:

# {
#   101: {
#     "unique_resources": 2,
#     "earliest_record": "2026-05-10T08:00:00Z"
#   },
#   202: {
#     "unique_resources": 2,
#     "earliest_record": "2026-05-11T09:00:00Z"
#   }
# }

from collections import defaultdict

logs = [
  {"batch_id": 101, "resource_id": "R1", "recorded_at": "2026-05-10T08:00:00Z"},
  {"batch_id": 101, "resource_id": "R2", "recorded_at": "2026-05-10T08:05:00Z"},
  {"batch_id": 101, "resource_id": "R1", "recorded_at": "2026-05-10T08:10:00Z"},
  {"batch_id": 202, "resource_id": "R1", "recorded_at": "2026-05-11T09:00:00Z"},
  {"batch_id": 202, "resource_id": "R3", "recorded_at": "2026-05-11T09:15:00Z"},
]

resources = defaultdict(set)
earliest_record = {}

for entry in logs:
    batch_id = entry["batch_id"]
    resource_id = entry["resource_id"]
    recorded_at = entry["recorded_at"]

    resources[batch_id].add(resource_id)

    if batch_id not in earliest_record or recorded_at < earliest_record[batch_id]:
        earliest_record[batch_id] = recorded_at

result = {}
for batch_id in resources:
    result[batch_id] = {
        "unique_resources": len(resources[batch_id]),
        "earliest_record": earliest_record[batch_id]
    }

print(result)
