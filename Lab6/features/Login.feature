Feature: Login

  Background:
    Given the following users exist:
      | Username | Password    |
      | testuser | testpassword|

  Scenario: Successful Login
    When logging in as user "testuser" with password "testpassword"
    Then the login MUST be successful

  Scenario: Wrong password
    When logging in as user "testuser" with password "wrongpassword"
    Then the login MUST NOT be successful
