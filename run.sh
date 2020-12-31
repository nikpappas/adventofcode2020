RE='^[0-9]+$'
if ! [[ $1 =~ $RE ]] ; then
   echo "Usage: ./run.sh <Number of day to run>" >&2; exit 1
else 
   go run src/*.go $1
fi
 


