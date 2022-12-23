# SSI Web Events

This is a simple web application that allows you to track events and their attendees.

under development.

##  Target Use Case

```html
<html>
  <body>
    <button type="button" data-we-name="register" data-we-listen="click">
      Register
    </button>
    <script type="text/javascript">
      window.ssiWe = window.ssiWe || {};
      window.ssiWe.config = {
        url: "https://ssi-we.yourdomain.com",
        token: "YOUR_TOKEN",
        user: {
          id: "USER_ID",
          name: "USER_NAME",
          email: "USER_EMAIL",
        }, // optional, if not provided, the user will be anonymous
      };
    </script>
    <script
      src="https://cdn.jsdelivr.net/gh/ssibrahimbas/ssi-we/client/ssi-we.js"
      type="text/javascript"
      defer
      async
    ></script>
  </body>
</html>
```
