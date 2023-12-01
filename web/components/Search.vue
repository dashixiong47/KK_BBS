<template>
  <div class="shadow-center w-2/6 h-10 flex items-center rounded-3xl px-3">
    <Icon class="w-6" name="ph:magnifying-glass-bold"></Icon>

    <input
      class="block w-full h-10 bg-inherit focus:border-none focus:outline-none"
      @keyup="enter"
      v-model="getQuery"
      placeholder="Start typing to search.."
      type="text"
    />
  </div>
</template>

<script setup>
const { getRouter } = usePath();
const router = useRouter();
const route = useRoute();
let searchValue = ref("");
function enter(e) {
  if (e.keyCode === 13) {
    if (searchValue.value) {
      router.push({
        path: getRouter("/search"),
        query: {
          q: searchValue.value,
        },
      });
    }
  }
}
const getQuery = computed({
  get: () => {
    if(route.path !== '/search') return ""
    return route.query.q;
  },
  set: (val) => {
    searchValue.value = val;
  },
});
</script>
