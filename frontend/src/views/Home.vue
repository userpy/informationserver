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
      axios.get('/api/user/contacts').catch(val => {
         console.log(`Ошибка =====${val}\n`)
      }).then(val => {
        console.log(` test axios ==========${JSON.stringify(val.data)}\n`)
      })
    }
  }
</script>
