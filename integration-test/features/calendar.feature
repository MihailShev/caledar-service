# file: features/calendar.feature

Feature: Create event
    Create event
    I need to be able to create some calendar event

    Scenario: Create event
        When I create calendar client
        Then I send message create event with params:
        """
        {
            "Title": "Test event",
            "Description": "test",
            "StartAfterNow": 5,
            "Duration": 5,
            "NotifyTime": -5,
            "UserId": 2
        }
        """
       And The response error should be empty

    Scenario: Get event
        When I send message get event
        Then I receive event with params:
        """
        {
           "UserId": 2,
           "Title": "Test event",
           "Description": "test"
        }
        """

    Scenario: Update event title
        When I send message update event with new title "Updated title"
        Then Event title match "Updated title"

    Scenario: Notification message is received
        When I received notify message
