<script setup>
import { ref } from 'vue'
import { normalizeTags } from '~/utils/normalizeTags'

const props = defineProps({
  id: { type: [String, Number], required: true },
  title: { type: String, required: true },
  description: { type: String, required: true },
  completed: { type: Boolean, default: false },
  tags: { type: Array, default: () => [] }
})

const emit = defineEmits(['updated'])

const { $api } = useNuxtApp()
const isAddingTag = ref(false)
const newTag = ref('')
const savingTag = ref(false)
const tagError = ref('')
const updatingComplete = ref(false)
const deletingTask = ref(false)

const openTagInput = () => {
  isAddingTag.value = true
  tagError.value = ''
}

const cancelTagInput = () => {
  if (savingTag.value) {
    return
  }

  isAddingTag.value = false
  newTag.value = ''
  tagError.value = ''
}

const addTag = async () => {
  const tags = normalizeTags(newTag.value)

  if (tags.length === 0) {
    tagError.value = 'Escribe al menos un tag antes de guardar.'
    return
  }

  savingTag.value = true
  tagError.value = ''

  try {
    await $api.tasks.addTagsToTask(props.id, tags)
    newTag.value = ''
    isAddingTag.value = false
    emit('updated')
  } catch (error) {
    console.error('Error al agregar el tag:', error)
    tagError.value = 'No se pudo agregar el tag.'
  } finally {
    savingTag.value = false
  }
}

const toggleComplete = async () => {
  if (props.completed) {
    return
  }

  updatingComplete.value = true

  try {
    await $api.tasks.completeTask(props.id)
    emit('updated')
  } catch (error) {
    console.error('Error al actualizar el estado de la tarea:', error)
  } finally {
    updatingComplete.value = false
  }
}

const deleteTask = async () => {
  const shouldDelete = window.confirm('¿Seguro que quieres eliminar esta tarea?')

  if (!shouldDelete) {
    return
  }

  deletingTask.value = true

  try {
    await $api.tasks.deleteTask(props.id)
    emit('updated')
  } catch (error) {
    console.error('Error al eliminar la tarea:', error)
  } finally {
    deletingTask.value = false
  }
}
</script>

<template>
  <div
    class="box mb-4 task-card"
    :style="{ backgroundColor: 'var(--bulma-card-background-color)', opacity: completed ? 0.9 : 1 }"
  >
    <div class="columns is-vcentered mb-0">
      <div class="column is-4">
        <p class="title is-5 mb-0" style="word-break: break-word;">
          {{ props.title }}
        </p>
      </div>

      <div class="column is-6">
        <p class="subtitle is-6 mb-0" style="word-break: break-word;">
          {{ props.description }}
        </p>
      </div>

      <div class="column is-2 has-text-right">
        <button
          class="button is-small"
          :class="[
            completed ? 'task-completed-button' : 'is-success',
            { 'is-loading': updatingComplete }
          ]"
          :disabled="completed || updatingComplete"
          @click="toggleComplete"
        >
          {{ completed ? 'Ya Completada' : '✔ Completar' }}
        </button>
      </div>
    </div>

    <hr class="my-2" style="background-color: var(--bulma-background); height: 1px;">

    <div class="columns is-vcentered mt-0">
      <div class="column is-4">
        <span class="has-text-weight-semibold is-size-7 mr-2">Estado:</span>
        <span class="tag task-status-tag" :class="completed ? 'task-status-completed' : 'task-status-pending'">
          {{ completed ? 'Completada' : 'Pendiente' }}
        </span>
      </div>

      <div class="column is-6">
        <div class="task-tags-header">
          <div class="task-tags-group">
            <div v-if="tags && tags.length > 0" class="tags mb-0 task-tags-list">
              <span class="has-text-weight-semibold is-size-7 mr-2 task-tags-label">Tags:</span>
              <span v-for="(tag, index) in tags" :key="index" class="tag task-tag-chip is-rounded">
                {{ tag }}
              </span>
            </div>
            <div v-else class="is-size-7 task-tags-empty">
              <span class="has-text-weight-semibold is-size-7 mr-2">Tags:</span> Sin etiquetas
            </div>

            <button class="button is-light task-add-tag-button" type="button" @click="openTagInput">
              +
            </button>
          </div>
        </div>

        <div v-if="isAddingTag" class="task-tag-form mt-3">
          <div class="field has-addons task-tag-addons">
            <div class="control is-expanded">
              <input
                v-model="newTag"
                class="input is-small"
                type="text"
                placeholder="Nueva tag"
                @keyup.enter="addTag"
                @keyup.esc="cancelTagInput"
              >
            </div>
            <div class="control">
              <button class="button is-small is-primary save-button" :class="{ 'is-loading': savingTag }" type="button" @click="addTag">
                Guardar
              </button>
            </div>
            <div class="control">
              <button class="button is-small cancel-button" type="button" :disabled="savingTag" @click="cancelTagInput">
                Cancelar
              </button>
            </div>
          </div>
          <p v-if="tagError" class="help is-danger">{{ tagError }}</p>
        </div>
      </div>

      <div class="column is-2 has-text-right">
        <button class="button is-danger is-outlined is-small" :class="{ 'is-loading': deletingTask }" :disabled="deletingTask" @click="deleteTask">
          🗑 Eliminar
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.task-card {
  border: 1px solid rgba(0, 0, 0, 0.06);
}

.task-tags-header {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  gap: 0.35rem;
}

.task-tags-group {
  display: flex;
  align-items: center;
  flex: 1;
  min-width: 0;
}

.task-tags-list,
.task-tags-empty {
  flex: 1;
}

.task-add-tag-button {
  flex: 0 0 auto;
  width: 1.55rem;
  height: 1.55rem;
  min-width: 1.55rem;
  min-height: 1.55rem;
  padding: 0;
  border-radius: 9999px;
  font-size: 0.9rem;
  line-height: 1;
  color: var(--bulma-text-strong);
  background-color: var(--bulma-white);
}

.task-status-tag,
.task-tag-chip {
  background-color: var(--bulma-white);
  color: var(--bulma-text-strong);
}

.task-status-completed,
.task-status-pending {
  background-color: var(--bulma-white);
  color: var(--bulma-text-strong);
}

.task-tag-form {
  max-width: 100%;
}

.task-tag-addons {
  display: flex;
  align-items: center;
  gap: 0.4rem;
}

.save-button {
  background-color: var(--bulma-background);
}

.cancel-button {
  background-color: var(--bulma-white);
}

.task-completed-button,
.task-completed-button[disabled] {
  background-color: #d7f2de;
  border-color: #d7f2de;
  color: #5c8f68;
  opacity: 1;
}

.task-completed-button:hover,
.task-completed-button[disabled]:hover {
  background-color: #d7f2de;
  border-color: #d7f2de;
  color: #5c8f68;
}
</style>
