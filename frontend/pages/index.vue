<script setup>
import { onMounted, watch } from 'vue'

const { refreshKey } = useTaskModal()
const {
  tareas,
  cargando,
  pag,
  limit,
  limitOptions,
  meta,
  obtenerTareas,
  cambiarLimit,
  cambiarPagina
} = usePaginatedTasks()

onMounted(() => {
  obtenerTareas()
})

watch(refreshKey, () => {
  obtenerTareas()
})
</script>

<template>
  <div>
    <section class="section">
      <div class="container">
        <div class="level mb-4" v-if="!cargando">
          <div class="level-left">
            <div class="level-item">
              <label for="page-size" class="mr-2">Mostrar</label>
              <div class="select is-small">
                <select id="page-size" :value="limit" @change="cambiarLimit">
                  <option v-for="opcion in limitOptions" :key="opcion" :value="opcion">
                    {{ opcion }}
                  </option>
                </select>
              </div>
              <span class="ml-2">items por página</span>
            </div>
          </div>
          <div class="level-right">
            <div class="level-item is-size-7">
              Total: {{ meta.totalItems }} tareas
            </div>
          </div>
        </div>
        
        <div v-if="cargando" class="has-text-centered my-6">
          <button class="button is-loading is-large loading-button">Cargando</button>
        </div>

        <div v-else-if="tareas.length === 0" class="notification empty-notification has-text-centered">
          No tienes tareas pendientes. ¡Buen trabajo! 🎉
        </div>

        <div v-else class="columns is-multiline">
          <div class="column is-12" v-for="tarea in tareas" :key="tarea.id">
            <TaskCard
                :id="tarea.id"
                :title="tarea.title" 
                :description="tarea.description" 
                :completed="tarea.completed"
                :tags="tarea.tags"
                @updated="obtenerTareas"
            />
          </div>
        </div>

        <nav
          v-if="!cargando && meta.totalPages > 1"
          class="pagination is-centered mt-5"
          role="navigation"
          aria-label="pagination"
        >
          <button
            class="pagination-previous"
            type="button"
            :disabled="Number(pag) <= 1"
            @click="cambiarPagina(Number(pag) - 1)"
          >
            Anterior
          </button>
          <button
            class="pagination-next pagination-next-button"
            type="button"
            :disabled="Number(pag) >= Number(meta.totalPages)"
            @click="cambiarPagina(Number(pag) + 1)"
          >
            Siguiente
          </button>

          <ul class="pagination-list">
            <li>
              <span class="pagination-link is-current pagination-current">
                Página {{ pag }} de {{ meta.totalPages }}
              </span>
            </li>
          </ul>
        </nav>
      </div>
    </section>
  </div>
</template>

<style scoped>
.pagination-previous {
  background-color: var(--bulma-box-background-color);
  color: var(--bulma-text-strong);
}

.pagination-next-button {
  background-color: var(--bulma-box-background-color);
  color: var(--bulma-text-strong);
}

.pagination-current {
  background-color: var(--bulma-dark);
  color: var(--bulma-white);
}
</style>