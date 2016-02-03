var URL_PATTERN = /^(https?:\/\/)?([a-zA-Z0-9\-]{1,63}\.)?([a-zA-Z0-9\-]{1,253}\.)([a-zA-Z0-9\-]{2,24}\.)?([a-zA-Z0-9\-]{2,24})(([/]{1})(.)*?)?$/;



function isValid(url) {
   return URL_PATTERN.test(url);
}

function testForURL(textBox) {
   var urlBoxText = textBox.value;
   var formDiv = document.getElementById('form-div');
   var valid = isValid(urlBoxText);

   if (valid) {
      formDiv.classList.add("has-success");
      formDiv.classList.remove("has-error");
   } else {
      formDiv.classList.add("has-error");
      formDiv.classList.remove("has-success");
   }

   return valid;
}

