<script setup>
import { RouterLink } from 'vue-router'
</script>

<template>
    <img class="img" src="../assets/stallion.png" alt="main-logo" />
    <small class="ref">
      Fast and Reliable Message Broker implemented with Go programming language.
      Get yourself on Stallion and deliver your events as fast as possible.
      Monitor your events, benchmark your system, set authentication for your clients and ...
    </small>
    <RouterLink class="start-btn" to="/about">
      Get Started!
    </RouterLink>
</template>

<style scoped>
.img {
  display: block;
  margin: 100px auto 0;
  inline-size: 40%;
  box-shadow: 0px 0px 10px 5px rgba(0, 0, 0, 0.2);
}

.ref {
  display: block;
  inline-size: 40%;
  margin: 50px auto;
  text-align: justify;
  color: rgb(120, 120, 120);
}

.start-btn {
  outline: none;
  text-decoration: none;
  color: #ffffff;
  background: rgb(255, 82, 82);
  padding: 10px 20px;
  border-radius: 5px;
  display: block;
  margin: 50px auto;
  inline-size: 140px;
}

.start-btn:hover {
  background: rgb(255, 55, 55);
  box-shadow: 0px 0px 4px 1px rgba(0, 0, 0, 0.2);
}
</style>