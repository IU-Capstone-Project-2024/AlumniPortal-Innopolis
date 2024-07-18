<template>
  <div class="p-4 mt-[6rem]">
    <h1 class="text-4xl font-bold text-center font-montserrat" style="color: #40BA21">Job offers</h1>
    <div class="mt-7 container flex flex-wrap gap-5 justify-around mx-auto">
      <UICard v-for="offer in offers" :key="offer.id" :image="offer.url" :header="offer.title"
              :description="offer.desc" :project="offer" @clicked="(id) => clickHandle(id)"/>
    </div>
  </div>
</template>

<script setup>
const router = useRouter();
const clickHandle = (id) => {
  console.log(id)
  //router.push('/offers/' + id)
}
const store = useOfferStore()

const offers = store.getOffers

</script>

<style lang="sass" scoped>

</style>