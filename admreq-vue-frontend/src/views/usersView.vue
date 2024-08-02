<template>
    <v-row>
        <v-col cols="6">
            <div class="pa-3">
                <v-list lines="one">
                    <v-list-item 
                        v-for="user in UsersStore.users"
                        :key="user.uuid"
                    >
                        <template v-slot:append>
                            <v-icon  color="blue" icon="mdi-lock" @click="pwdid = user.uuid, UsersStore.passDialog = true"></v-icon>
                            <v-icon  color="red" icon="mdi-delete-forever" @click="delid = user.uuid, dialog = true"></v-icon>

                        </template>
                        <v-list-item-title>
                            {{ user.lastname }} {{ user.firstname }} {{ user.surname }}
                        </v-list-item-title>
                        <v-list-item-subtitle>
                            {{ user.departmentName}} / 
                            <v-icon slot="activator" class="mr-1 ml-1" v-if="user.userRights.create" icon="mdi-card-text-outline" color="teal-darken-1" label="создание"></v-icon>
                            <v-icon class="mr-1 ml-1" v-if="!user.userRights.create" icon="mdi-card-text-outline"></v-icon>
                            <v-icon class="mr-1 ml-1" v-if="user.userRights.users" icon="mdi-account-multiple" color="teal-darken-1" size="small"></v-icon>
                            <v-icon class="mr-1 ml-1" v-if="!user.userRights.users" icon="mdi-account-multiple"></v-icon>
                            <v-icon class="mr-1 ml-1" v-if="user.userRights.employee" icon="mdi-pipe-wrench" color="teal-darken-1"></v-icon>
                            <v-icon class="mr-1 ml-1" v-if="!user.userRights.employee" icon="mdi-pipe-wrench"></v-icon>
                            <v-icon class="mr-1 ml-1" v-if="user.userRights.admin" icon="mdi-sitemap" color="teal-darken-1"></v-icon>
                            <v-icon class="mr-1 ml-1" v-if="!user.userRights.admin" icon="mdi-sitemap"></v-icon>
                            <v-icon class="mr-1 ml-1" v-if="user.userRights.stat" icon="mdi-archive" color="teal-darken-1"></v-icon>
                            <v-icon class="mr-1 ml-1" v-if="!user.userRights.stat" icon="mdi-archive"></v-icon>
                            <v-icon class="mr-1 ml-1" v-if="user.userRights.archiv" icon="mdi-matrix" color="teal-darken-1"></v-icon>
                            <v-icon class="mr-1 ml-1" v-if="!user.userRights.archiv" icon="mdi-matrix"></v-icon>
                        </v-list-item-subtitle>

                    </v-list-item>
                </v-list>
                <v-dialog v-model="dialog" width="500">
                    <v-card
                        prepend-icon="mdi-trash-can"
                        text="Вы уверены, что хотите заблокировать пользователя? Доступ к системе обращений будет прекращен!"
                        title="Блокировка пользователя"
                    >
                    <template v-slot:actions>
                    <v-spacer></v-spacer>

                <v-btn fab dark small color="primary" @click="dialog = false">
                Отменить
                </v-btn>

                <v-btn fab dark small color="red" @click="UsersStore.deleteUser(delid, AuthStore.credentials.token), dialog = false">
                    заблокировать
                </v-btn>
            </template>
            </v-card>
        </v-dialog>

        <v-dialog v-model="UsersStore.passDialog" width="500">
                    <v-card
                        prepend-icon="mdi-form-textbox-password"
                        title="Смена пароля пользователя"
                    >
                    <v-form @submit.prevent="UsersStore.addDepartment(depart, dowork, AuthStore.credentials.token)">
                        <v-text-field v-model="passwd" label="Новый пароль"></v-text-field>
                    </v-form>
                    <template v-slot:actions>
                    <v-spacer></v-spacer>

                <v-btn fab dark small color="primary" @click="UsersStore.passDialog = false">
                Отменить
                </v-btn>

                <v-btn fab dark small color="red" @click="UsersStore.changeUserPassword(pwdid, passwd, AuthStore.credentials.token)">
                    Сменить
                </v-btn>
            </template>
            </v-card>
        </v-dialog>
            </div>
        </v-col>
        <v-col cols="3">
            <v-card class="pa-3 ma-3" variant="elevated" elevation="16" color="teal-lighten-4">
                <span class="ma-2">Добавить отдел:</span><br><br>
            <v-form @submit.prevent="UsersStore.addDepartment(depart, dowork, AuthStore.credentials.token)">
                <v-text-field v-model="depart" label="Название отдела"></v-text-field>
                <v-switch label="Работает с обращениями" color="teal-darken-1" v-model="dowork"></v-switch>
                <v-btn type="submit" color="teal-darken-1" block class="mt-2">Добавить</v-btn>
            </v-form>
            </v-card>
        </v-col>
        <v-col cols="3">
                <v-card class="pa-3 ma-3" variant="elevated" elevation="16" color="teal-lighten-4">
                    <span class="ma-2">Добавить пользователя:</span><br><br>
                <v-form fast-fail @submit.prevent="UsersStore.createUser(firstname, lastname, surname, department, username, password, rights, AuthStore.credentials.token)">
                    <v-text-field v-model="lastname" label="Фамилия"></v-text-field>
                    <v-text-field v-model="firstname" label="Имя"></v-text-field>
                    <v-text-field v-model="surname" label="Отчество"></v-text-field>
                    <v-select
                        label="Отдел"
                        v-model="department"
                        :items="UsersStore.departments"
                        item-title="department"
                        item-value="uuid"
                    >
                    </v-select>
                    <v-text-field v-model="username" label="Логин"></v-text-field>
                    <v-text-field v-model="password" label="Пароль"></v-text-field>
                    <v-select
                        label="Права доступа"
                        multiple
                        v-model="rights"
                        :items="[{title: 'Создание', value: 'create'}, {title: 'Исполнение', value: 'employee'}, {title: 'Управление', value: 'admin'}, {title: 'Пользователи', value: 'users'}, {title: 'Архив', value: 'archiv'}, {title: 'Статистика', value: 'stat'}]"
                    ></v-select>
                    <v-btn type="submit" color="teal-darken-1" block class="mt-2">Добавить</v-btn>
                </v-form>
                <span class="d-flex align-center justify-center ma-3 text-red">{{ UsersStore.usersErrors }}</span>
                </v-card>
        </v-col>
    </v-row>
</template>

<script setup>
import { useAuthStore } from '../stores/AuthStore';
import { useUsersStore } from '../stores/UsersStore';
import { ref } from 'vue';

const AuthStore = useAuthStore();
const UsersStore = useUsersStore();

UsersStore.getUsers(AuthStore.credentials.token);
UsersStore.getDepartments(AuthStore.credentials.token)

const firstname = ref("");
const lastname = ref("");
const surname = ref("");
const department = ref("");
const username = ref("");
const password = ref("");
const rights = ref();
const dowork = ref(false);
const depart = ref("");
const dialog = ref(false);
const passwd = ref("");
var delid = "";
var pwdid = "";
</script>