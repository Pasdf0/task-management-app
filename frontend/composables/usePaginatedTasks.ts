import { ref } from 'vue'

type PaginationMeta = {
  totalItems: number
  totalPages: number
  currentPage: number
  pageSize: number
}

export const usePaginatedTasks = () => {
  const { $api } = useNuxtApp()

  const tareas = ref<any[]>([])
  const cargando = ref(true)
  const pag = ref(1)
  const limit = ref(10)
  const limitOptions = [5, 10, 20, 50]
  const meta = ref<PaginationMeta>({
    totalItems: 0,
    totalPages: 1,
    currentPage: 1,
    pageSize: 10
  })

  const obtenerTareas = async () => {
    cargando.value = true

    try {
      const respuesta = await $api.tasks.getSomeTasks({
        page: pag.value,
        limit: limit.value
      })

      const payload = respuesta.data || {}
      tareas.value = Array.isArray(payload.data) ? payload.data : []

      meta.value = {
        totalItems: Number(payload.meta?.totalItems || 0),
        totalPages: Number(payload.meta?.totalPages || 1),
        currentPage: Number(payload.meta?.currentPage || pag.value),
        pageSize: Number(payload.meta?.pageSize || limit.value)
      }

      pag.value = meta.value.currentPage
    } catch (error) {
      console.error('Error al cargar tareas:', error)
      tareas.value = []
    } finally {
      cargando.value = false
    }
  }

  const cambiarLimit = async (event: Event) => {
    const target = event.target as HTMLSelectElement | null
    limit.value = Number(target?.value || limit.value)
    pag.value = 1
    await obtenerTareas()
  }

  const cambiarPagina = async (nuevaPagina: number) => {
    const currentPage = Number(pag.value)
    const totalPages = Number(meta.value.totalPages)
    const nextPage = Number(nuevaPagina)

    if (!Number.isFinite(nextPage)) {
      return
    }

    if (nextPage < 1 || nextPage > totalPages || nextPage === currentPage) {
      return
    }

    pag.value = nextPage
    await obtenerTareas()
  }

  return {
    tareas,
    cargando,
    pag,
    limit,
    limitOptions,
    meta,
    obtenerTareas,
    cambiarLimit,
    cambiarPagina
  }
}
