// component for header

import PropTypes from 'prop-types'

const Header = ({ headerTitle }) => {
    return (
        <div>
            <h1>{headerTitle}</h1>
        </div>
    )
}


Header.defaultProps = {
    //default header title
    headerTitle: 'Upcoming',
}

Header.propTypes = {
    headerTitle: PropTypes.string.isRequired,
}

export default Header
