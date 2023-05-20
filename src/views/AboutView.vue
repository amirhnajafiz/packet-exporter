<template>
  <div>
    <h1>
      About Stallion
    </h1>
    <p class="about-text">
      Stallion is a Message Broker that is implemented with Go programming language.
      You can use Stallion to publish events on specific topics and consume them over client side with Stallion SDKs.
      Stallion provides authentication so you can authorise your clients.<br />
      One good thing about Stallion is that it uses only Golang internal modules and has no external modules.
      Everything in Stallion is pure Golang.
      It comes with Load Test (Race) and Blackbox Exporter (Vet) to analyse the Stallion server.
    </p>
  </div>
</template>

<style scoped>
.about-text {
  margin: 50px 0 200px;
  text-align: justify;
}
</style>