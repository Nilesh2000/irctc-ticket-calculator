## Problem statement

We want to develop a module to help IRCTC to calculate the ticket price for a given train. Idea is to give this module as an extension to different apps in future. Below are some examples

## Example 1

A train number 12345 has only sleeper and general coaches. The train runs between Mumbai to Pune and it has stops as shown below

| Start - Mumbai | Stop 1 - Karjat | Stop 2 - Lonavala | Stop 3 - Chinchwad | Stop 4 - Pune |
|----------------|-----------------|-------------------|--------------------|---------------|

Passengers can board at any stop and leave at _any stop ahead._

**This train follows fix-per-station pricing strategy as below**

For each stop passenger has to pay

Rs. 20/- for general coach

Rs. 40/- for sleeper coach

### Cases

| \# of passengers | Train number | Coach   | Start station | End station | Total Price              |
|------------------|--------------|---------|---------------|-------------|--------------------------|
| 1                | 12345        | general | Karjat        | Chinchwad   | **40**/-                 |
| 3                | 12345        | sleeper | Mumbai        | Pune        | 3 \* 40 \* 4 = **480**/- |
