<template>
  <div>
    {{data}}
  </div>
</template>

<script>
  import axios from 'axios'
  export default {
    name: 'Home',
    data:function () {
        return{
          data:'test'
        }
    },
    components: {
    },
    created() {
      axios.get('/api/user/contacts').catch(val => {
         console.log(`Ошибка =====${val}\n`)
      }).then(val => {
        console.log(` test axios ==========${JSON.stringify(val)}\n`);
        this.data = val.data;
      })
    }
  }
</script>
