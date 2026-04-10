<script setup>
const { successMessage, clearSuccess } = useTaskModal()
</script>

<template>
  <div>
    <NavBar />
    <TaskCreateModal />

    <div v-if="successMessage" class="app-toast-wrap">
      <article class="notification is-success app-toast">
        <button class="delete" @click="clearSuccess"></button>
        {{ successMessage }}
      </article>
    </div>
    
    <main>
      <slot />
    </main>
  </div>
</template>

<style scoped>
.app-toast-wrap {
  position: fixed;
  top: 1rem;
  right: 1rem;
  z-index: 60;
}

.app-toast {
  min-width: 260px;
  max-width: 420px;
}
</style>