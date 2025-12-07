usage() {
    echo """
    Simple command-line script for initializing a new daily AoC challenge with a Go module setup.

    USAGE:
    bash generate_day.sh --day/-d DAY_NUMBER
    
    REQUIRED OPTIONS:
    d | day: number of the day to generate the project for
    
    WARNING: you need to have go installed for this to work!
    """
    exit 1
}

day_num=""

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

new_dir_name="day-${day_num}"

mkdir $new_dir_name
cd $new_dir_name
go mod init $new_dir_name
echo "package main" > main.go
echo "package main" > main_test.go
touch README.md
touch input.txt
touch test.txt
