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
                        <v-text-field
                                :rules="passwordConfirmRules"
                                v-model="passwordConfirm"
                                label="Подтвердить пароль"
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
                        <v-text-field
                                v-model="email"
                                :rules="emailRules"
                                label="Email"
                                type="email"
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
                        <v-btn small  >
                            Регистрация
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
        name: "Registration",
        props:{
            formType: String,
        },
        methods:{
            state: function(val){
                this.$emit('update:formType', val);
            },
            Password(v) {
                if (this.password == v) {
                    return true;
                } else {
                    return false;
                }
            }
        },
        data:function () {
            return{
                passwordRules:[
                    v => !!v || 'Пароль обязательное поле',
                    v => v.length >= 8 || 'Пароль должен быть больше 8 символов',
                ],
                passwordConfirmRules:[
                    v => !!v || 'Пароль обязательное поле',
                    v => this.Password(v) || 'Пароли должны совпадать'
                ],
                nameRules: [
                    v => !!v || 'Name обязательное поле',
                    v => v.length >= 6 || 'Имя должно быть больше 6 символов',
                ],
                emailRules: [
                    v => !!v || 'E-mail обязательное поле',
                    v => /.+@.+/.test(v) || 'E-mail должен быть корректным',
                ],
                login:null,
                password: null,
                passwordConfirm: null,
                email: null
            }
        },
    }
</script>

<style scoped>

</style>