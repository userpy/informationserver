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
                                outlined
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
                                :rules="passwordRules"
                                v-model="password"
                                label="Пароль"
                                type="password"
                                required
                                outlined
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
           getLength(val, len) {
               try{
                   return  val.length >= len ? true: false
               }catch (e) {
                   return null
               }
           }
        },
        name: "Login",
        data:function () {
            return{
                nameRules:[
                    v => !!v || "Это поле не должно быть пустым",
                    v => this.getLength(v, 4)|| 'Имя должно быть больше 4 символов',
                ],
                passwordRules:[
                    v => !!v || "Это поле не должно быть пустым",
                    v => this.getLength(v, 8) || 'Пароль должен быть больше 8 символов',
                ],
                login:null,
                password: null
            }
        },
    }
</script>

<style scoped>

</style>