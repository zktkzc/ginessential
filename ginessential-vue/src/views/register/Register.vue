<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="用户注册">
          <b-form>
            <b-form-group label="用户名">
              <b-form-input
                  v-model="$v.user.name.$model"
                  type="text"
                  placeholder="请输入用户名（选填）"
              ></b-form-input>
            </b-form-group>
            <b-form-group label="手机号">
              <b-form-input
                  v-model="$v.user.telephone.$model"
                  type="number"
                  placeholder="请输入手机号"
                  :state="validateState('telephone')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('telephone')">
                手机号不符合要求
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group label="密码">
              <b-form-input
                  v-model="$v.user.password.$model"
                  type="password"
                  placeholder="请输入密码"
                  :state="validateState('password')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('password')">
                密码至少6位
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group>
              <b-button variant="outline-primary" block @click="register"
              >注册
              </b-button
              >
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import {required, minLength} from "vuelidate/lib/validators";
import customValidator from "@/helper/validator";
import {mapActions} from "vuex";

export default {
  data() {
    return {
      user: {
        name: "",
        telephone: "",
        password: "",
      },
    };
  },
  validations: {
    user: {
      name: {},
      telephone: {
        required,
        telephone: customValidator.telephoneValidator,
      },
      password: {
        required,
        minLength: minLength(6),
      },
    },
  },
  methods: {
    ...mapActions('userModule', {userRegister: 'register'}),
    validateState(name) {
      const {$dirty, $error} = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      // 验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }

      // 请求API
      this.userRegister(this.user).then(() => {
        // 跳转主页
        this.$router.replace({name: "Home"});
      }).catch(err => {
        this.$bvToast.toast(err.response.data.msg, {
          title: "数据验证错误",
          variant: "danger",
          solid: true,
        });
      });
    },
  },
};
</script>

<style lang="scss" scoped>
</style>