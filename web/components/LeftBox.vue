<template>
  <div>
    <CardTitleCard
      title="More Pages"
      v-model="actived"
      :list="data"
      class="mb-5 max-h-80"
    />
    <!-- <CardTitleCard title="More Pages" class="mb-5 h-80" /> -->
  </div>
</template>

<script setup>
import { useGroupStore } from "~/stores/init.js";
const store = useGroupStore();
const data = computed(() => store.getGroup);
const actived = computed({
  get: () => store.actived,
  set: (val) => {
    store.setActived(val);
  },
});

store.fetchGroup();
</script>
