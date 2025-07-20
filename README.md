# Cronicle

Cronicle is a tool that converts cryptic cron expressions into clear, human-readable text.

## Minute
- **Input:** `1` → At 1 minute past the hour
- **Input:** `1,2` → At 1 and 2 minutes past the hour
- **Input:** `1-8` → Minutes 1 through 8 past the hour
- **Input:** `*` → Every minute
- **Input:** `4/9` → Every 9 minutes, starting at 4 minutes past the hour

## Hour
- **Input:** `1` → between 1:00 AM and 1:00 AM (need to fix this)
- **Input:** `1,2,13` → At 1:00 AM, 2:00 AM, and 1:00 PM
- **Input:** `0/2` → Every 2 hours, starting at 12:00 AM
- **Input:** `9-17` → between 9:00 AM and 5:00 PM (need to fix this)

## Day of Month
- **Input:** `1` → On day 1 of the month
- **Input:** `1,2,13` → On days 1, 2, and 13 of the month
- **Input:** `1/2` → Every 2 days
- **Input:** `3/5` → Every 5 days, starting on day 3 of the month
- **Input:** `9-17` → Between day 9 and 17 of the month

## Month
- **Input:** `1` → Only in January
- **Input:** `1,2,11` → Only in January, February, and November
- **Input:** `1/3` → Every 3 months
- **Input:** `4/2` → Every 2 months, from April through December
- **Input:** `5-8` → May through August
- **Input:** `JAN,MAR` → Only in January and March
- **Input:** `JUL-OCT` → July through October

## Day of Week
- **Input:** `MON-WED` → Monday through Wednesday
- **Input:** `1-5` → Monday through Friday
- **Input:** `1,3,5` → Only on Monday, Wednesday, and Friday
- **Input:** `0` → Only on Sunday