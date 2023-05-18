<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <div>
      <nav>
        <RouterLink class="link" to="/">Home</RouterLink>
        <RouterLink class="link" to="/about">About</RouterLink>
      </nav>
    </div>
  </header>
  <RouterView />
</template>

<style scoped>
nav {
  inline-size: 100%;
  margin: 20px 0;
  padding: 0 20px;
}

.link {
  color: rgb(80, 80, 80);
  outline: none;
  text-decoration: none;
  margin-inline-end: 30px;
}

.link:hover {
  color: rgb(120, 120, 120);
}
</style>
