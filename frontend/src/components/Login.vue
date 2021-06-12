<template>
    <v-app id="inspire">
        <v-form  class="fa-border">
            <v-container >
                <v-row>
                    <v-col>
                    </v-col>
                    <v-col cols="12"
                           md="3">
                            <v-btn-toggle
                                    v-model="text"
                                    tile
                                    color="deep-purple accent-3"
                                    group
                            >
                            <v-btn value="left" @click="state('login')" >
                                Вход
                            </v-btn>

                            <v-btn value="center" @click="state('reg')">
                                Регистрация
                            </v-btn>
                        </v-btn-toggle>
                    </v-col>
                    <v-col>
                    </v-col>
                </v-row>
                <v-row >
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
                                v-model="password"
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
                        <v-btn small @click="state('login')">
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
    export default {
        props:{
            formType: String,
        },
        methods:{
           state: function(val){
              this.$emit('update:formType', val);
          },
        },
        name: "Login",
        data:function () {
            return{
                login:null,
                password: null
            }
        },
    }
</script>

<style scoped>

</style>