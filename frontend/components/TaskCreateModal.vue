<script setup>
import { ref } from 'vue'
import { normalizeTags } from '~/utils/normalizeTags'

const { $api } = useNuxtApp()
const { isOpen, close, bumpRefresh, showSuccess } = useTaskModal()

const title = ref('')
const description = ref('')
const tagsInput = ref('')
const saving = ref(false)
const errorMessage = ref('')

const resetForm = () => {
  title.value = ''
  description.value = ''
  tagsInput.value = ''
  errorMessage.value = ''
}

const handleClose = () => {
  if (saving.value) {
    return
  }

  close()
  resetForm()
}

const closeAfterSuccess = () => {
  close()
  resetForm()
}

const submitTask = async () => {
  if (!title.value.trim()) {
    errorMessage.value = 'El titulo es obligatorio.'
    return
  }

  saving.value = true
  errorMessage.value = ''

  try {
    const tags = normalizeTags(tagsInput.value)

    await $api.tasks.createTask({
      title: title.value.trim(),
      description: description.value.trim() || 'Sin descripcion',
      tags
    })

    bumpRefresh()
    closeAfterSuccess()
    showSuccess('Tarea guardada correctamente.')
  } catch (error) {
    console.error('Error al crear la tarea:', error)
    errorMessage.value = 'No se pudo crear la tarea. Intenta de nuevo.'
  } finally {
    saving.value = false
  }
}
</script>

<template>
  <div v-if="isOpen" class="modal is-active" role="dialog" aria-modal="true" aria-label="Crear nueva tarea">
    <div class="modal-background" @click="handleClose"></div>

    <div class="modal-card">
      <header class="modal-card-head task-modal-head">
        <p class="modal-card-title">Nueva Tarea</p>
        <button class="delete" aria-label="close" @click="handleClose"></button>
      </header>

      <section class="modal-card-body">
        <div class="field">
          <label class="label" for="task-title">Titulo</label>
          <div class="control">
            <input
              id="task-title"
              v-model="title"
              class="input"
              type="text"
              maxlength="120"
              placeholder="Ej: Preparar reporte semanal"
            >
          </div>
        </div>

        <div class="field">
          <label class="label" for="task-description">Descripcion</label>
          <div class="control">
            <textarea
              id="task-description"
              v-model="description"
              class="textarea"
              rows="4"
              maxlength="500"
              placeholder="Detalles de la tarea"
            ></textarea>
          </div>
        </div>

        <div class="field">
          <label class="label" for="task-tags">Tags</label>
          <div class="control">
            <input
              id="task-tags"
              v-model="tagsInput"
              class="input"
              type="text"
              placeholder="Ej: urgente, trabajo, personal"
            >
          </div>
          <p class="help">Separa cada tag con coma.</p>
        </div>

        <p v-if="errorMessage" class="help is-danger">{{ errorMessage }}</p>
      </section>

      <footer class="modal-card-foot">
        <button class="button is-primary task-save-button" :class="{ 'is-loading': saving }" @click="submitTask">
          Guardar
        </button>
        <button class="button task-cancel-button" :disabled="saving" @click="handleClose">Cancelar</button>
      </footer>
    </div>
  </div>
</template>

<style scoped>
.modal-card {
  width: min(92vw, 640px);
}

.task-modal-head {
  background-color: var(--bulma-card-background-color);
}

.task-save-button {
  background-color: var(--bulma-background);
}

.task-cancel-button {
  background-color: var(--bulma-white);
}
</style>
