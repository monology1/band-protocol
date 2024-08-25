
# Band-Protocol Assignment

## Running the “Boss Baby” Problem
1. Navigate to the baby-boss-revenge directory:
```bash
cd band-protocol/baby-boss-revenge
```
2. Run the program:
```bash
go run main.go
```

3. Expected Output:
   This Program will run test cases and print test case passed or failed based on the expected output.

## Running the “Superman Chicken Rescue” Problem
1.	Navigate to the superman-chicken-rescue directory:
```bash
cd band-protocol/superman-chicken-rescue
```
2.	Run the program:
```bash
go run main.go
```
3.	Expected Output:
      This Program will run test cases for both the brute-force and optimized solutions, and printing test case passed or failed.

## Running the “Transaction Broadcasting and Monitoring” Problem

1.	Navigate to the transaction-broadcasting-and-monitoring-client directory:
```bash
cd band-protocol/transaction-broadcasting-and-monitoring-client
```
2. Install dependencies:
```bash
go mod tidy
```
3. Run the application:
```bash
go run main.go
```
4.	Broadcasting a Transaction:
      Send a POST request to the /broadcast endpoint
```bash
curl -X POST http://localhost:8080/broadcast \
-H "Content-Type: application/json" \
-d '{
    "symbol": "ETH",
    "price": 4500
}'
```
5.	Monitoring the Transaction Status:
      The transaction status will be monitored automatically after broadcasting. The system will continue to check the status until it is CONFIRMED, FAILED, or DNE.

## Explanation of Functions
Boss Baby Revenge Problem
- isBossBabyGood: Determines if Boss Baby is a “Good boy” or a “Bad boy” based on the sequence of shots (S) and revenge moves (R). The logic checks if all shots (S) have been avenged by corresponding revenge moves (R).

Superman Chicken Rescue Problem
- BruteForceSupermanChickenRescue: Implements a brute-force solution to determine the maximum number of chickens that can be rescued within a given distance k. Time complexity is O(n^2).
- OptimizeSupermanChickenRescue: Implements an optimized solution using a sliding window technique. Time complexity is O(n log n) due to sorting.

Transaction Broadcasting and Monitoring Problem
- BroadcastTransaction(c Context) (string, error): Handles broadcasting a transaction, including generating the timestamp internally.
- MonitorTransaction(txHash string) (string, error): Periodically checks the transaction status until it is either CONFIRMED, FAILED, or DNE.
- BroadcastAndMonitorTransaction(c Context): Combines the broadcasting and monitoring functions to provide an end-to-end solution.

## Error Handling

- Broadcasting Errors: If broadcasting fails (e.g., due to network issues), the system returns a 500 Internal Server Error with a relevant message.
- Monitoring Errors: If the system fails to retrieve the transaction status, it returns a 500 Internal Server Error.
- Status Handling: The system processes transaction statuses (CONFIRMED, PENDING, FAILED, DNE) explicitly, providing clear feedback on the transaction’s state.
