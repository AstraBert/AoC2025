usage() {
    echo """
    Simple command-line script for getting the solution from an AoC challenge built in Go.

    USAGE:
    bash get_solution.sh --day/-d DAY_NUMBER --complex
    
    REQUIRED OPTIONS:
    d | day: number of the day to get the solution for
    complex: get the solution for part 2. Defaults to false (get the solution for part 1). 
    
    WARNING: you need to have go installed and 'input.txt' filled with your input data for this to work!
    """
    exit 1
}

day_num=""
complex="false"

# Loop through the commandline arguments
while
    [[ $# -gt 0 ]] \
        ;
do
    case "$1" in
    -h | --help)
        usage
        ;;
    -d | --day)
        day_num="$2"
        shift 2
        ;;
    --complex)
        complex="true"
        shift 1
        ;;
    *)
        echo "Unknown option: $1"
        usage
        ;;
    esac
done

if [[ -z $day_num ]]
then 
    echo "Missing required argument: '-d/--day'"
    usage
fi

dir_name="day-${day_num}"

if [[ ! -d $dir_name ]]
then
    echo "No such directory: ${dir_name}"
    exit 2
fi

cd $dir_name
if [[ $complex == "false" ]]
then
    echo "Getting the solution for part 1 of day ${day_num}"
    go run main.go simple
else
    echo "Getting the solution for part 2 of day ${day_num}"
    go run main.go complex
fi