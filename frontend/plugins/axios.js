import axios from 'axios'

export default defineNuxtPlugin(() => {
	const config = useRuntimeConfig()

	const apiBaseUrl = import.meta.server
		? (config.apiInternalBaseUrl || config.public.apiBaseUrl)
		: config.public.apiBaseUrl
	const apiTimeout = Number(config.public.apiTimeout || 10000)

	if (!apiBaseUrl) {
		throw new Error('API base URL is not configured. Set NUXT_PUBLIC_API_BASE_URL and optionally NUXT_API_INTERNAL_BASE_URL.')
	}

	const apiClient = axios.create({
		baseURL: apiBaseUrl,
		timeout: apiTimeout,
		headers: {
			'Content-Type': 'application/json'
		}
	})

	const tasks = {
		createTask: (task) => apiClient.post('/tasks', task),
		getAllTasks: (params = {}) => apiClient.get('/tasks', { params }),
		getSomeTasks: (params = {}) => apiClient.get('/tasks/filter', { params }),
		getTask: (id) => apiClient.get(`/tasks/${id}`),
		completeTask: (id, payload = {}) => apiClient.put(`/tasks/${id}/complete`, payload),
		deleteTask: (id) => apiClient.delete(`/tasks/${id}`),
		addTagsToTask: (id, tags = []) => apiClient.post(`/tasks/${id}/tags`, { tags })
	}

	return {
		provide: {
			axios: apiClient,
			api: {
				tasks
			}
		}
	}
})
