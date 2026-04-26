## Lab 1:

Goal: Prove you can process files larger than the available RAM without crashing the system (OOM).

Generate a "Monster" File: Create a script (or use a tool) to generate a 5GB CSV file with random user data (ID, Name, Email, Timestamp).

Implement a Buffered Reader: Use Go's bufio.Scanner or csv.Reader to process the file line-by-line rather than loading the whole file into a slice.

Data Transformation: Convert the timestamp to ISO-8601 and lowercase all emails.

The "Dry Run" Output: Print a progress report every 100,000 rows showing current memory usage using runtime.ReadMemStats.

Validation: Demonstrate that your memory usage stays "flat" (e.g., under 20MB) even while processing a multi-gigabyte file.

-------------------------------------------
## Lab 2:

### Goal:  Show mastery of Go Routines and Channels to bypass the "one-at-a-time" bottleneck of network calls.

Define the Sources: Identify 3 different public APIs (or mock local HTTP servers) that provide related data (e.g., Weather, Traffic, and Local Events).

Worker Pool Pattern: Spin up a fixed number of "Worker" goroutines (e.g., 5 workers).

The "Job" Queue: Feed the URLs into a channel. The workers pick up a URL, make the request, and send the result to a "Results" channel.

Synchronization: Use a sync.WaitGroup to ensure the program doesn't exit until all workers have finished their tasks.

Fan-In Result: Collect all results into a single final "Report" map and print the total execution time compared to a sequential version.

------------------

## Lab 3

### Goal: Tackle the "Space Complexity" problem we discussed by handling millions of unique keys efficiently.

Input Stream: Simulate a stream of 1 million "Transaction IDs" where many are duplicates.

The Pre-allocated Map: Based on your knowledge of Go map internals, initialize your tracking map with an explicit size hint to prevent "doubling" triggers.

The Filter: Logic to check if an ID exists. If not, add it to the map and write it to an "Output" file.

Memory Benchmarking: Run the lab twice—once with make(map[string]bool) and once with the size hint.

Report: Record the time difference and peak memory usage between the two versions to show you understand Go’s allocation overhead.

----------------------

## Lab 4

### Goal: Demonstrate "Production-Ready" code that doesn't lose data when a server restarts or a cloud function times out.

Setup a Pipeline: Create a process that reads from a simulated "Queue" (a channel) and writes to a simulated "Database" (a local file).

Context Integration: Use context.WithCancel to wrap your main processing loop.

Signal Listening: Use os/signal to listen for a SIGINT (Ctrl+C).

The "Flush" Logic: When a signal is received, stop accepting new "Jobs" but allow the current workers 5 seconds to finish writing their data to the file before exiting.

Evidence: Show logs that prove the program finished its last write after the user hit Ctrl+C but before the program closed.

---------------------

## Lab 5

### Goal: Show how you handle "Dirty Data" in an industrial pipeline.

Source "Messy" JSON: Create a file with 1,000 JSON objects. Some are perfect, some have missing fields, and some have wrong data types (e.g., a string where an int should be).

Strict Unmarshalling: Use Go Struct tags (json:"field_name") and custom validation logic to check each record.

The Branching Path:

Success: If valid, send to the "Clean" channel.

Failure: If invalid, send to a "Dead-Letter" channel with an attached error message.

Multi-File Output: Write the "Clean" data to success.json and the "Dirty" data to errors.log.

Summary Statistics: Output a final count (e.g., "Processed 1000: 950 Success, 50 Failed").