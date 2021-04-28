<template>
  <hello-world />
</template>

<script>
  import HelloWorld from '../components/HelloWorld'
  import axios from 'axios'
  export default {
    name: 'Home',

    components: {
      HelloWorld,
    },
    created() {
      axios.get('https://mocki.io/v1/d4867d8b-b5d5-4a48-a4ab-79131b5809b8').catch(val => {
         console.log(`Ошибка =====${val}\n`)
      }).then(val => {
        console.log(` test axios ==========${JSON.stringify(val)}\n`)
      })
    }
  }
</script>
