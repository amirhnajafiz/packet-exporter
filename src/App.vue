<script setup>
import Navbar from './components/Navbar.vue';
import { RouterView } from 'vue-router'
</script>

<template>
  <header>
    <Navbar />
  </header>
  <RouterView class="wrapper" />
</template>

<style scoped>
.wrapper {
  margin: 100px auto;
  inline-size: 60%;
  color: rgb(120, 120, 120);
}
</style>
