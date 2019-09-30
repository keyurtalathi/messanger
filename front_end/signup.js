var app = angular.module('app', ['toaster']);

app.controller("SignUpController", function($scope, $window, $http, toaster){

    var reset = function(){
        $scope.name="";
        $scope.email="";
        $scope.alias="";
        $scope.contact="";
        $scope.password="";
    }
    var reset1 = function(){
        $scope.loginemail="";
        $scope.loginpassword="";
    }
    $scope.sign_up = function(){
        var name = $scope.name;
        var email = $scope.email;
        var alias = $scope.alias;
        var contact = $scope.contact;
        var password = $scope.password;
        if(name!="" && email!="" && alias!="" 
            && contact!="contact" && password!=""){
            $http({
                method : "POST",
                url : "http://127.0.0.1/api/signup",
                header:{
                "Content-Type":"application/json",
                "Access-Control-Allow-Origin": "127.0.0.1:1234"
                },
                data:{   
                    "Username" : name,
                    "Password" : password,
                    "Alias"    : alias,
                    "Phone"    : contact,
                    "Email"    : email
                }
            }).then(function mySuccess(response) {
                console.log(response)
                toaster.pop('success', "success", "Registration successful");
                reset();
            },function error(response){
                toaster.pop('error', "error", "Incorrect data");
            });
    }
    else{
        toaster.pop('error', "error", "Fill up data");
    }
    }

    
    $scope.login = function(){
        var email = $scope.loginemail;
        var password = $scope.loginpassword;
            $http({
                method : "POST",
                url : "http://127.0.0.1/api/authenticate",
                header:{
                  "Content-Type":"application/json",
                  "Access-Control-Allow-Origin": "127.0.0.1:1234"
                },
                data:{   
                    "Email" : email,
                    "Password" : password
                }
              }).then(function mySuccess(response) {
                  console.log(response)
                  var response = response.data.responseData
                  localStorage.setItem("email",email)
                  localStorage.setItem("token",response.Token)
                  toaster.pop('success', "success", "Login successful");
                  reset1();
                  $window.location.href = "messanger.html"
              },function error(response){
                $scope.errorMsg="Incorrect username or password"
                alert("invalid")
            });
        }
});