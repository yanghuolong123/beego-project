{{if .user}}
<div class="alert alert-success" role="alert">
 你已经登陆, 用户名:{{.user.Username}}, 邮箱: {{.user.Email}}。 <a class="btn btn-success" href="logout">退出</a>
</div>
{{else}}
<form method="post">
  <div class="form-group">
    <label for="exampleInputEmail1">邮件:</label>
    <input type="email" name="email" class="form-control" id="exampleInputEmail1" placeholder="Email">
  </div>
  <div class="form-group">
    <label for="exampleInputPassword1">密码:</label>
    <input type="password" name="password" class="form-control" id="exampleInputPassword1" placeholder="Password">
  </div>
  <div class="checkbox">
    <label>
      <input type="checkbox"> 记住密码 
    </label>
  </div>
  <button type="submit" class="btn btn-default">提交</button>
</form>
{{end}}
