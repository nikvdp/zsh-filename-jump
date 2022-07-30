function zsh_word_back_widget() {
  local res=$(zsh-back-word "$BUFFER" "$CURSOR" "b")
  CURSOR=$((res - 1))
}

function zsh_word_forward_widget() {
  local res=$(zsh-back-word "$BUFFER" "$CURSOR" "w")
  CURSOR=$((res + 1))
}

# register them as zsh widgets
zle -N zsh_word_back_widget 
zle -N zsh_word_forward_widget

# and bind them to B and W respectively
bindkey -M vicmd W zsh_word_forward_widget
bindkey -M vicmd B zsh_word_back_widget
