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
import { useGetGroupDetail } from "~/api/server";
const store = useGroupStore();
const data = computed(() => store.getGroup);
const list = [0, 1, 2, 3, 4, 0, 1, 2, 3, 4];
let groupActived = ref(0);
store.fetchGroup();
</script>
