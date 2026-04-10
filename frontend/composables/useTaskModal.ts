let successTimer: ReturnType<typeof setTimeout> | null = null

export const useTaskModal = () => {
  const isOpen = useState<boolean>('task-modal-open', () => false)
  const refreshKey = useState<number>('task-list-refresh-key', () => 0)
  const successMessage = useState<string>('task-success-message', () => '')

  const open = () => {
    isOpen.value = true
  }

  const close = () => {
    isOpen.value = false
  }

  const bumpRefresh = () => {
    refreshKey.value += 1
  }

  const showSuccess = (message: string, duration = 2500) => {
    successMessage.value = message

    if (successTimer) {
      clearTimeout(successTimer)
    }

    successTimer = setTimeout(() => {
      successMessage.value = ''
      successTimer = null
    }, duration)
  }

  const clearSuccess = () => {
    successMessage.value = ''

    if (successTimer) {
      clearTimeout(successTimer)
      successTimer = null
    }
  }

  return {
    isOpen,
    refreshKey,
    successMessage,
    open,
    close,
    bumpRefresh,
    showSuccess,
    clearSuccess
  }
}
