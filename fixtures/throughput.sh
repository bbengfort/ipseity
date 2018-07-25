#!/bin/bash
# Runs the throughput benchmark

# Location of the results
RESULTS="throughput-0.3-local.csv"

RUNS=12
MIN_CLIENTS=1
MAX_CLIENTS=12

# Describe the time format
TIMEFORMAT="experiment completed in %2lR"

time {
  # Write header to the output file
  # echo "server,clients,messages,duration,throughput" >> $RESULTS


  for (( I=0; I<=$RUNS; I+=1 )); do
      # Step Four: Run benchmarks with 3-6 clients
      for (( J=$MIN_CLIENTS; J<=$MAX_CLIENTS; J++ )); do

        if [ $J -lt 36 ]; then
          UPTIME=5s
        else
          UPTIME=10s
        fi

        # for SERVER in simple sequence actor locker ; do
        #
        #   ipseity serve -u $UPTIME -s $SERVER &
        #   sleep 1
        #   ipseity bench -s $SERVER -c $J >> $RESULTS &
        #   wait
        #
        # done

        ipseity serve -u $UPTIME -s stream &
        sleep 1
        ipseity bench -s stream -c $J >> $RESULTS &
        wait

      done
  done
}
