selection=$(hummingbird)

if [[ "$selection" == "x" ]]; then
    # Do nothing if 'x' is returned
    true
elif [[ -d "$selection" ]]; then
    # If it's a directory, change to it
    cd "$selection"
elif [[ -f "$selection" ]]; then
    # If it's a file, first get its directory
    dir_of_file=$(dirname "$selection")
    filename=$(basename "$selection") # Get just the filename for nvim

    # Change to the directory of the file
    if [[ -d "$dir_of_file" ]]; then # Check if dirname returned a valid directory
        cd "$dir_of_file"
    else
        # This case should be rare if hummingbird returns a valid file path
        echo "Warning: Could not determine directory for '$selection'."
        # Script will continue, nvim will be called with full path if it happens
    fi

    # Now, check if the file (now referenced by its basename relative to the new CWD)
    # is executable.
    if [[ -x "$filename" ]]; then
        # If the file is executable, do nothing (don't open it with nvim)
        # The cd to its directory has already happened.
        # echo "Selected file '$filename' is executable, not opening."
        true
    else
        # If the file is not executable, open it with neovim
        # Since we've cd'd into its directory, we can just use the filename
        nv "$filename"
    fi
else
    # Handle cases where selection is not "x", not a directory, and not a regular file
    if [[ -n "$selection" ]]; then # Check if selection is not empty
        echo "Hummingbird returned: '$selection'. Not a recognized file or directory."
    fi
fi

clear
# ls -hl will show the contents of the current directory.
# This will now always be the directory of the selected item (or the original if 'x')
ls -hl
