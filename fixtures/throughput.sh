#!/bin/bash
# Runs the throughput benchmark

# Location of the results
RESULTS="throughput-1.0-local.csv"

RUNS=12
MIN_CLIENTS=21
MAX_CLIENTS=42

# Describe the time format
TIMEFORMAT="experiment completed in %2lR"

time {
  # Write header to the output file
  # echo "clients,messages,duration,throughput" >> $RESULTS


  for (( I=0; I<=$RUNS; I+=1 )); do
      # Step Four: Run benchmarks with 3-6 clients
      for (( J=$MIN_CLIENTS; J<=$MAX_CLIENTS; J++ )); do

        ipseity serve -u 5s &
        sleep 1
        ipseity bench -c $J >> $RESULTS &
        wait

      done
  done
}
