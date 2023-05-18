<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <div>
      <nav>
        <RouterLink class="link" to="/">Home</RouterLink>
        <RouterLink class="link" to="/about">About</RouterLink>
        <RouterLink class="link" to="/">Documents</RouterLink>
        <RouterLink class="link right-link" to="/">Github</RouterLink>
      </nav>
    </div>
  </header>
  <RouterView class="wrapper" />
</template>

<style scoped>
.wrapper {
  margin: 40px auto;
  inline-size: 60%;
}

nav {
  display: flex;
  align-content: center;
  font-size: 0.8em;
  inline-size: 100%;
  padding: 30px 100px;
  background: rgb(255, 109, 109);
}

.link {
  background: #ffffff;
  padding: 7px;
  border-radius: 5px;
  color: rgb(30, 30, 30);
  outline: none;
  text-decoration: none;
  margin-inline-end: 50px;
}

.link:hover {
  color: rgb(120, 120, 120);
}

.right-link {
  margin-inline-start: auto;
  order: 2;
}
</style>
