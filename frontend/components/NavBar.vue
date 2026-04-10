<script setup>
const { open } = useTaskModal()
</script>

<template>
  <nav class="navbar is-primary app-navbar" role="navigation" aria-label="main navigation">
    <div class="container">
      <div class="navbar-brand">
        <NuxtLink to="/" class="navbar-item brand-link">
          <h1 class="title is-4 mb-0 brand-title">Task Management App</h1>
        </NuxtLink>

        <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false">
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>

      <div class="navbar-menu">
        <div class="navbar-start">
          <NuxtLink to="/" class="navbar-item nav-link">
            Inicio
          </NuxtLink>
        </div>

        <div class="navbar-end">
          <div class="navbar-item">
            <div class="buttons">
              <button class="button is-light add-task-button"
              type="button" @click="open">
                <strong>+ Nueva Tarea</strong>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<style scoped>
.app-navbar {
  background-color: var(--bulma-box-background-color);
}

.brand-title,
.brand-link,
.nav-link {
  color: var(--bulma-text-strong);
}

.app-navbar .navbar-item:hover,
.app-navbar .navbar-item:focus {
  background-color: rgba(68, 55, 66, 0.08);
  color: var(--bulma-text-strong);
}

.app-navbar .navbar-burger span {
  background-color: var(--bulma-text-strong);
}

.add-task-button {
  background-color: var(--bulma-box-background-color);
  color: var(--bulma-text-strong);
}
</style>