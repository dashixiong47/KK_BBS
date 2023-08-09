<template>
  <div class="col-span-8 bg-slate-200">
    <h1 class="text-2xl font-bold">User Info</h1>
    <p>id: {{ $route.params.id }}</p>
  </div>
</template>

<script setup>
definePageMeta({
  layout: "userInfo",
});
</script>
