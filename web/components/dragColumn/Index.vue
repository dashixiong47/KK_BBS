<script setup>
import Sortable from "sortablejs";

const myList = ref(null);
let { items } = defineProps({
  items: {
    type: Array,
    default: () => [],
  },
});

onMounted(() => {
  new Sortable(myList.value, {
    animation: 150,
    ghostClass: "sortable-ghost",
    onUpdate: function (evt) {
      const oldIndex = evt.oldIndex;
      const newIndex = evt.newIndex;
      const movedItem = items.value.splice(oldIndex, 1)[0];
      items.value.splice(newIndex, 0, movedItem);
    },
  });
});
</script>

<template>
  <div>
    <ul ref="myList">
      <li v-for="(item, index) in items" :key="item" class="flex border">
        <div class="min-w-[50px]">{{ index }}</div>
        <div>{{ item.name }}</div>
      </li>
    </ul>
  </div>
</template>
