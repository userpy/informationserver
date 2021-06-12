<template>
    <v-app id="inspire">
        <v-form  class="fa-border">
            <v-container >
                <v-row>
                    <v-col>

                    </v-col>
                    <v-col
                            cols="12"
                            md="3"
                    >
                        <v-text-field
                                v-model="firstname"
                                :rules="nameRules"
                                label="Логин"
                                required
                        ></v-text-field>
                    </v-col>
                    <v-col>

                    </v-col>
                </v-row>
                <v-row>
                    <v-col>

                    </v-col>
                    <v-col
                            cols="12"
                            md="3"
                    >
                        <v-text-field
                                v-model="lastname"
                                label="Пароль"
                                type="password"
                                required
                        ></v-text-field>
                    </v-col>
                    <v-col>

                    </v-col>
                </v-row>
                <v-row>
                    <v-col>

                    </v-col>
                    <v-col
                            cols="12"
                            md="3"
                    >
                        <v-btn>
                            Войти
                        </v-btn>
                    </v-col>
                    <v-col>

                    </v-col>
                </v-row>
            </v-container>
        </v-form>
    </v-app>
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
