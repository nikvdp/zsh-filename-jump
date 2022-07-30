function zsh_word_back_widget() {
  local res=$(zsh-back-word "$BUFFER" "$CURSOR" "b")
  CURSOR=$((res - 1))
}

function zvm_after_lazy_keybindings() {
  zvm_define_widget zsh_word_back_widget

  zvm_bindkey vicmd 'B' zsh_word_back_widget
}
