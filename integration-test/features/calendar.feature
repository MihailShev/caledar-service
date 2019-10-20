# file: features/calendar.feature

Feature: Calendar service
    I need to be able to create some calendar event and get notify
    And I need to able to get created event and update it

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
       And I received notify message with created event

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

