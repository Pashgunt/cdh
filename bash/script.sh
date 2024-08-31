function cd() {
    builtin cd "$@" && go-dir add "$(pwd)"
}

function cdh() {
    local dir
    dir=$(go-dir get "$1")
    if [ -d "$dir" ]; then
        builtin cd "$dir" && go-dir add "$dir"
    else
        echo "The directory does not exist: $dir"
    fi
}

function cdi() {
    local dirs
    dirs=($(go-dir list | awk '{print $2}'))

    if [ ${#dirs[@]} -eq 0 ]; then
        echo "The directory history is empty."
        return
    fi

    if command -v fzf > /dev/null 2>&1; then
        local selection
        selection=$(go-dir list | fzf --height 40% --reverse | awk '{print $2}')
        if [ -n "$selection" ] && [ -d "$selection" ]; then
            builtin cd "$selection" && go-dir add "$selection"
        else
            echo "Incorrect selection or the directory does not exist."
        fi
    else
        echo "Select the directory to navigate to:"
        local i=1
        for dir in "${dirs[@]}"; do
            echo "$i: $dir"
            ((i++))
        done

        local choice
        if [[ -n "$ZSH_VERSION" ]]; then
            read "choice?Enter the directory number: "
        else
            read -p "Enter the directory number: " choice
        fi

        if [[ $choice =~ ^[0-9]+$ ]] && [ $choice -ge 1 ] && [ $choice -le ${#dirs[@]} ]; then
            local selected_dir=${dirs[$((choice))]}
            if [ -d "$selected_dir" ]; then
                builtin cd "$selected_dir" && go-dir add "$selected_dir"
            else
                echo "The directory does not exist: $selected_dir"
            fi
        else
            echo "Wrong choice."
        fi
    fi
}