var app = angular.module('app', ['toaster']);
app.controller("MessangerController", function($scope, $rootScope ,$http, toaster,fileReader){
    $scope.openChat=null
    $scope.create_web_socket = function(){
      ws = new WebSocket("ws://localhost:1234/echo?token="+localStorage.getItem("token"));
        ws.onopen = function(evt) {
            // alert("Welcome.......");
        }
        ws.onclose = function(evt) {
            alert("Byyyyeeeeee");
            ws = null;
        }
        ws.onmessage = function(evt) {
          var msg=evt.data.split(":")
          //   console.log("**************88")
          // console.log($scope.openChat.ContactList[0])
          for(var i=0;i<$scope.Chats.length;i++){
            if(msg[0]!="group"){
              if($scope.Chats[i].ContactList[0].ContactEmail==msg[1])
              { 

                var obj={
                  "MessageBody" : msg[5],
                  "MessageType":msg[3],
                  "Reciever":[msg[2]],
                  "Sender":msg[1],
                  "Status":msg[4],
                  "Self" : false
                }
                $scope.Chats[i].MessageList.push(obj);
                console.log("$scope.Chats[i].MessageList")
                console.log($scope.Chats[i].MessageList)
                if($scope.openChat!=null){
                  document.getElementById($scope.openChat.id).click()
                }
                break;
              }
            }
            else{
              if($scope.Chats[i].GroupId==msg[2])
              { 

                var obj={
                  "MessageBody" : msg[5],
                  "MessageType":msg[3],
                  "Reciever":[msg[2]],
                  "Sender":msg[1],
                  "Status":msg[4],
                  "Self" : false
                }
                $scope.Chats[i].MessageList.push(obj);
                console.log("$scope.Chats[i].MessageList")
                console.log($scope.Chats[i].MessageList)
                if($scope.openChat!=null){
                  document.getElementById($scope.openChat.id).click()
                }
                break;
              }
            }
          }
        
        }
        ws.onerror = function(evt) {
            alert("ERROR: " + evt.data);
        }
    }()


    $scope.imageSrc = "";
    
    $scope.$on("fileProgress", function(e, progress) {
      $scope.progress = progress.loaded / progress.total;
    });
    var reset = function(){
      $scope.groupName="";
      $scope.contactName="";
      $scope.contactEmail="";
    }
    $scope.display_chat_show=true;
    $scope.select_contact_show=false;
    $scope.new_group_show=false;
    $scope.create_group_show=false;
    $scope.show_chat=false;
    $scope.open_group_data = false;

    $scope.select_contact = function(){
      $scope.select_contact_show=true;
      $scope.display_chat_show=false;
    }
    $scope.set_div =function(){
      if($scope.new_group_show==true){
        $scope.select_contact_show=true;
        $scope.new_group_show=false;
      }
      else if($scope.select_contact_show==true){
        $scope.select_contact_show=false;
        $scope.display_chat_show=true;
      }
      else if($scope.create_group_show==true){
        $scope.create_group_show=false;
        $scope.new_group_show=true;
      }
      else if($scope.show_chat==true){
        $scope.show_chat=false;
      }
      else{
        $scope.display_chat_show=true;
      }
    }
    $scope.edit_group = function(){
      $scope.open_group_data = true;
      $scope.show_chat=false;
    }
    $scope.close_group = function(){
      $scope.open_group_data = false;
      $scope.show_chat=true;
    }

    $scope.new_group = function(){
      var members = [];
      // console.log("contacts")
      // console.log($scope.Contacts)
      reset();
      for(var i=0;i<$scope.Contacts.length;i++){
        members.push($scope.Contacts[i]["ContactList"][0]);
        members[i].Selected = false;
      }
      // console.log("members")
      // console.log(members)
      var block_list = $scope.UserContact.ContactObject.BlockList;
      for(var i=0;i<members.length;i++){
        for(var j=0;j<block_list.length;j++){
          if(members[i].ContactEmail==block_list[j].ContactEmail){
            members.splice(i,1);
            break; 
          }
        }
      }
      // console.log("mmmenee")
      // console.log(members)
      $scope.Members = members;
      reset();
      // console.log($scope.Members)
      $scope.new_group_show=true;
      $scope.select_contact_show=false;
    }
    $scope.create_group = function(){
      var list =[]
        for (var i = 0; i < $scope.Members.length; i++) {
            if ($scope.Members[i].Selected) {
                var obj = $scope.Members[i];
                list.push(obj)
            }
        }
        $scope.selectedList = list;
        if($scope.selectedList.length == 0){
          toaster.pop('error', "error", "Add atleast one contact");
        }
        else{
          $scope.new_group_show=false;
          $scope.create_group_show=true;
        }
    }
    $scope.create_group_data = function(){
      var emailList = []
      var groupname = $scope.groupName
      var image =""
      var token = localStorage.getItem("token")
      for (var i = 0; i < $scope.selectedList.length; i++) {
        emailList.push($scope.selectedList[i].ContactEmail)
      }
      if(groupname!=null){
        $http({
          method : "POST",
          url : "http://127.0.0.1/api/group",
          header:{
            "Content-Type":"application/json",
            "Access-Control-Allow-Origin": "127.0.0.1:1234"
          },
          data:{
            "Token":token,
            "GroupMembers": emailList ,
            "GroupName" : groupname,
            "GroupImage": $rootScope.url
          }
        }).then(function mySuccess(response) {
            // console.log(response)
            toaster.pop('success', "success", "Group created successfully");
            $scope.create_group_show=false;
            $scope.display_chat_show=true;
        });
      }
      else{
        toaster.pop('error', "error", "Provide a group name");

      }
      }
    $scope.add_contact = function(){
        var contactEmail = $scope.contactEmail
        var contactName = $scope.contactName
        var token = localStorage.getItem("token");
        // console.log("email"+contactEmail+"name"+contactName)
        $http({
          method : "PUT",
          url : "http://127.0.0.1/api/contactlist",
          header:{
            "Content-Type":"application/json",
            "Access-Control-Allow-Origin": "*"
          },
          data:{
            "Token": token,
            "ContactPayload": {
              "ContactEmail" : contactEmail,
              "ContactName" : contactName
            },
            "Type" : "addcontact"
          }
        }).then(function mySuccess(response) {
            $('#modalwindow').modal('hide');
            reset();
        });
      }
      $scope.remove_contact = function(contact){
        var contactEmail = contact.ContactList[0].ContactEmail;
        var contactName = contact.ContactList[0].ContactName;
        var token = localStorage.getItem("token");
        $http({
          method : "PUT",
          url : "http://127.0.0.1/api/contactlist",
          header:{
            "Content-Type":"application/json",
            "Access-Control-Allow-Origin": "*"
          },
          data:{
            "Token": token,
            "ContactPayload": {
              "ContactEmail" : contactEmail,
              "ContactName" : contactName
            },
            "Type":"removecontact"
          }
        }).then(function mySuccess(response) {
            toaster.pop('success', "success", "Contact removed successfully");
            // console.log(response)
        });
      }
      $scope.block_contact = function(contact){
        var contactEmail = contact.ContactList[0].ContactEmail;
        var contactName = contact.ContactList[0].ContactName;
        var token = localStorage.getItem("token");
        $http({
          method : "PUT",
          url : "http://127.0.0.1/api/contactlist",
          header:{
            "Content-Type":"application/json",
            "Access-Control-Allow-Origin": "*"
          },
          data:{
            "Token": token,
            "ContactPayload": {
              "ContactEmail" : contactEmail,
              "ContactName" : contactName
            },
            "Type":"blockcontact"
          }
        }).then(function mySuccess(response) {
          $scope.Blocked = true;
          toaster.pop('success', "success", "Contact blocked successfully");
        });
      }


      $scope.unblock_contact = function(contact){
        var contactEmail = contact.ContactList[0].ContactEmail;
        var contactName = contact.ContactList[0].ContactName;
        var token = localStorage.getItem("token");
        $http({
          method : "PUT",
          url : "http://127.0.0.1/api/contactlist",
          header:{
            "Content-Type":"application/json",
            "Access-Control-Allow-Origin": "*"
          },
          data:{
            "Token": token,
            "ContactPayload": {
              "ContactEmail" : contactEmail,
              "ContactName" : contactName
            },
            "Type":"unblock"
          }
        }).then(function mySuccess(response) {
          toaster.pop('success', "success", "Contact unblocked successfully");
          $scope.start();
          $scope.Blocked=false;
        });
      }

      $scope.open_chat = function(chat){
        // console.log(chat)
        $scope.Blocked = false;
        var block_list = $scope.UserContact.ContactObject.BlockList;
        for(var i=0;i<block_list.length;i++){
          if(chat.ContactList[0].ContactEmail==block_list[i].ContactEmail){
            $scope.Blocked = true;
            break;
          }
        }
        
        $scope.openChat = chat;
        $scope.open_group_data = false;
        $scope.show_chat = true;
        $scope.MessageCopyList = chat.MessageList;
        console.log("$scope.MessageCopyList");
        console.log($scope.MessageCopyList);
        var myEmail = $scope.UserContact.UserObject.Email
        for(var i=0;i< $scope.MessageCopyList.length;i++){
          if($scope.MessageCopyList[i].Sender==myEmail)
          {
            $scope.MessageCopyList[i].Self=true;
          }
          else
          {
            $scope.MessageCopyList[i].Self=false;
          }
        }
      }
      $scope.init = function(){
        var email = localStorage.getItem("email");
        var token = localStorage.getItem("token"); 
        $http({
          method : "GET",
          url : "http://127.0.0.1/api/chat",
          header:{
              "Content-Type":"application/json"
          },
          params:{
              "email" : email
          }
        }).then(function mySuccess(response) {
          $scope.Flag=0;
          $scope.ChatData=response.data.responseData
          $scope.Chats=[]
            //construct object for displaying chats
            for (k=0;k<$scope.ChatData.length;k++){
              var displayChat = {}
              
              if($scope.ChatData[k].MessageList.length!=0 || $scope.ChatData[k].GroupId!="" ){
                var list=$scope.ChatData[k].ContactList;
                var userlist=$scope.UserContact.ContactObject.ContactList;
              displayChat.Email = $scope.ChatData[k].Email;
              displayChat.GroupId = $scope.ChatData[k].GroupId;
              displayChat.GroupName = $scope.ChatData[k].GroupName;
              displayChat.GroupImage = $scope.ChatData[k].GroupImage;
              displayChat.MessageList = $scope.ChatData[k].MessageList;
              displayChat.RecentMessage = $scope.ChatData[k].RecentMessage;
              displayChat.id="chat"+k;
              if($scope.ChatData[k].GroupId==""){
                displayChat.Flag=false;
              }
              else{
                displayChat.Flag=true;
              }
              for(var j=0;j<list.length;j++){
                var saved=0;
                for(var i=0;i<userlist.length;i++){
                  if(list[j].ContactEmail==userlist[i].ContactEmail){
                    list[j].ContactName = userlist[i].ContactName;
                    saved=1;
                    break;
                  }
                }
                if(saved==0){
                  list[j].ContactName = list[j].ContactEmail;
                }
              }
              displayChat.ContactList=list;
              $scope.Chats.push(displayChat);
              console.log("chat list")
              console.log($scope.Chats)
            }
            }
            //construct object to select contact
          $scope.Contacts=[];
          var contactlist = $scope.UserContact.ContactObject.ContactList;
          console.log("----contact------");
          console.log(contactlist);
          for(var i=0;i<$scope.ChatData.length;i++){
            var obj = {}
            if($scope.ChatData[i].GroupId==""){
              for(var j=0;j<contactlist.length;j++){
                if($scope.ChatData[i].ContactList[0].ContactEmail==contactlist[j].ContactEmail){
                  obj = $scope.ChatData[i]
                  obj.ContactList[0].ContactName = contactlist[j].ContactName;
                  $scope.Contacts.push(obj);
                }
              }
            }
          }
          console.log("----------");
          console.log($scope.Contacts);
        });
        }
      $scope.start = function(){
      $scope.Flag=0;
        var token = localStorage.getItem("token")
        $http({
          method : "GET",
          url : "http://127.0.0.1/api/authenticate",
          header:{
              "Content-Type":"application/json"
          },
          params:{
              "token" : token
          }
        }).then(function mySuccess(response) {
          $scope.UserContact=response.data.responseData;
          $scope.Flag=1;
        });
      }
      $scope.send_message = function(open_chat){
        // console.log(open_chat)
        var obj={
          "MessageType":"msg",
          "Reciever":[open_chat.ContactList[0].ContactEmail],
          "Sender":open_chat.Email,
          "Status":"sent",
          "MessageBody" : $scope.message,
          "Self" : true
        }
        $scope.MessageCopyList.push(obj)
        console.log($scope.MessageCopyList)
        if(open_chat.GroupId=="")
        { 
          ws.send("personal:"+open_chat.Email+":"+open_chat.ContactList[0].ContactEmail+":msg:sent:"+$scope.message);
        }
        else{

          ws.send("group:"+open_chat.Email+":"+open_chat.GroupId+":msg:sent:"+$scope.message);
        }
        $scope.message=""
      }

      window.setInterval(function(){
        if($scope.Flag==1)
        {
          $scope.init();
        }
      }, 1000);

      

      $scope.file_upload = function(){
        document.getElementById('my_file').click();
      }
      $scope.start();
});

  app.directive("ngFileSelect", function(fileReader, $timeout,$http, $rootScope) {
    return {
      scope: {
        ngModel: '='
      },
      link: function($scope, el) {
        function getFile(file) {
          fileReader.readAsDataUrl(file, $scope)
            .then(function(result) {
              $timeout(function() {
                $scope.ngModel = result;
              });
            });
        }

        el.bind("change", function(e) {
          var file = (e.srcElement || e.target).files[0];
          getFile(file);
          // console.log(file)
          var formData = new FormData()
          formData.append("file", file)
          $http.post("http://127.0.0.1:2428/upload", formData, {
              transformRequest: angular.identity,
              headers: {
                'Content-Type': undefined,
            }
          }).then(function mySuccess(response) {
              $rootScope.url = response.data.url;
              if($rootScope.url!=null){
                alert("Image Uploaded Successfully");
              }
          });
          // $http({
          //   method : "POST",
          //   url : "http://127.0.0.1/api/upload",
          //   header:{
          //     "Content-Type":"application/json",
          //     "Access-Control-Allow-Origin": "*"
          //   },
          //   data:{
          //     "File":formData
          //   }
          // }).then(function mySuccess(response) {
          //   console.log("file upload")
          //     console.log(response)
          //     $rootScope.url = response.data.url;
          // });
        });
      }
    };
  });

app.factory("fileReader", function($q, $log) {
  var onLoad = function(reader, deferred, scope) {
    return function() {
      scope.$apply(function() {
        deferred.resolve(reader.result);
      });
    };
    // console.log("00000")
    // console.log(reader.result)
  };

  var onError = function(reader, deferred, scope) {
    return function() {
      scope.$apply(function() {
        deferred.reject(reader.result);
      });
    };
  };

  var onProgress = function(reader, scope) {
    return function(event) {
      scope.$broadcast("fileProgress", {
        total: event.total,
        loaded: event.loaded
      });
    };
  };

  var getReader = function(deferred, scope) {
    var reader = new FileReader();
    reader.onload = onLoad(reader, deferred, scope);
    reader.onerror = onError(reader, deferred, scope);
    reader.onprogress = onProgress(reader, scope);
    return reader;
  };

  var readAsDataURL = function(file, scope) {
    var deferred = $q.defer();

    var reader = getReader(deferred, scope);
    reader.readAsDataURL(file);

    return deferred.promise;
  };

  return {
    readAsDataUrl: readAsDataURL
  };
});
