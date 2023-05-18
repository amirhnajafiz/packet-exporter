<script setup>
import Footer from './components/Footer.vue';
import Navbar from './components/Navbar.vue';
import { RouterView } from 'vue-router'
</script>

<template>
  <header>
    <Navbar />
  </header>
  <div class="wrapper">
    <RouterView />
  </div>
  <footer>
    <Footer />
  </footer>
</template>

<style scoped>
.wrapper {
  margin: 150px auto;
  inline-size: 90%;
  color: rgb(120, 120, 120);
}
</style>
