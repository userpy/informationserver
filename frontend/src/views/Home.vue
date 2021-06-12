<template>
<div>
    <login v-if="formType == 'login'" :form-type.sync="formType"/>
    <registration v-else-if="formType == 'reg'" :form-type.sync="formType" />

</div>

</template>

<script>
  import Registration from "../components/Registration";
  import Login from "../components/Login";
  import axios from 'axios'
  export default {
    name: 'Home',
    data:function () {
        return{
            formType: 'login',
        }
    },
    components: {
        Login, Registration
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
