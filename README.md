# Simple web go application

This application uses **[gin](https://github.com/gin-gonic/gin)** and is building with standart tools for linux host, which allow it to run as a standalone application.
To build use the custom **[go builder](https://hub.docker.com/layers/159441296/privetkakdela/devops_course/go_builder/images/sha256-46e2588e18895400c81d42bb1138e77df1835ed3b3a5a93e4e47a58994c9a14e?context=repo)** with all the required libraries installed, which makes the builder as standalone.
And the app is packaging into a docker container.

# How use?

Clone repo, install **```docker```** and run **```docker build -t app_go .```** Run docker conteiner **```docker run -d -p 8080:8080 --name app_go app_go```**
Check result **```http://yours_ip:8080```**

### Result:
![app_go](/app_two/pic/res.png)