# Requirements Document

## Introduction

This feature involves extracting the mortgage rate display logic from the FirstTimeBuyer page into a reusable MortgageRateCard component with clean, professional styling. The component will display individual mortgage rate information in a visually appealing card format.

## Glossary

- **MortgageRateCard**: A React component that displays mortgage rate information for a single lender
- **FirstTimeBuyer**: The existing page component that displays first-time buyer mortgage rates
- **Rate Data**: Object containing lender information including product, rate, APRC, max LTV, and monthly payment

## Requirements

### Requirement 1

**User Story:** As a developer, I want to extract the mortgage rate display logic into a reusable component, so that I can maintain cleaner code and reuse the component across different pages.

#### Acceptance Criteria

1. THE MortgageRateCard SHALL accept rate data as props
2. THE MortgageRateCard SHALL display all mortgage rate information fields (lender, product, rate, APRC, max LTV, monthly payment)
3. THE MortgageRateCard SHALL be located in the components directory
4. THE FirstTimeBuyer SHALL use the new MortgageRateCard component instead of inline div elements
5. THE MortgageRateCard SHALL maintain the same data display functionality as the original implementation

### Requirement 2

**User Story:** As a user, I want the mortgage rate information to be displayed in a visually appealing card format, so that I can easily read and compare different mortgage options.

#### Acceptance Criteria

1. THE MortgageRateCard SHALL have a clean, card-like visual design with proper spacing
2. THE MortgageRateCard SHALL use consistent typography and color scheme
3. THE MortgageRateCard SHALL have proper visual hierarchy with the lender name prominently displayed
4. THE MortgageRateCard SHALL have hover effects for better user interaction
5. THE MortgageRateCard SHALL be responsive and work well on different screen sizes
