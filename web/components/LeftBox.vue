<template>
  <div>
    <CardTitleCard
      title="More Pages"
      v-model="groupActived"
      :list="data"
      class="mb-5 max-h-80"
    />
    <CardTitleCard title="More Pages" class="mb-5 h-80" />
  </div>
</template>

<script setup>
import { useGroupStore } from "~/stores/init.js";
const store = useGroupStore();
const data = computed(() => store.getGroup);
let groupActived = ref(0);
store.fetchGroup();
</script>
