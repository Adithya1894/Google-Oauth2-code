# google-oauth-code

1. Procedure for installing Golang can be found here: https://golang.org/doc/install/source.

2. Clone this project into your local repository where Gopath is set.

3. Edit the client id and client secret key in the lines 34 and 35 of logincontroller/logincontroller.go.

4. Use command go run DraftSolution.go to build and run the project.

5. Open http://127.0.0.1:9090/login in any of the browsers to see the working project, where It enables you to login with your google credentials.

# Manually Creating a Dataset/Table in bigquery.
 
1. The file Bqdataset.go allows you to manually create dataset and table in Google Bigquery. 

2. You can use the documentation in https://developers.google.com/ad-exchange/rtb/open-bidder/google-app-guide to create your own project.

3. Edit the lines 16,25 and 40 in Bqdataset.go to update names of your project, dataset and table.

4. Use command go run Bqdataset.go to build and run this project.

5. In your Bigquery console, changes in dataset, table can be seen.



