import { render, fireEvent } from '@testing-library/react'
import WTag from './tag'
import { AllInclusive } from '@mui/icons-material'
import './index.scss'

describe('WTag', () => {
    it('Should redirect to pages correctly as specified', () => {
        const path = '/ruta-especifica'
        const loggedMessages = []

        //Redirect the output of console.log to loggedMessages
        const originalLog = console.log
        console.log = (...args) => {
            loggedMessages.push(args.join(' '))
        };

        const { getByText } = render(<WTag icon={AllInclusive} text={path} path={path} />); //Chage "text" to "path"
        fireEvent.click(getByText(path)); //Change "TagLink" to "path"

        //Restore the original console.log function
        console.log = originalLog

        //Verify that it has been called with the expected message
        expect(loggedMessages).toContain(`Ir a la ruta ${path}`);
    })

    it('Should display style variations when isActive is true or false', () => {
        const { container } = render(<WTag icon={AllInclusive} text="TagLink" isActive={true} />);
        expect(container.firstChild).toHaveClass('tagLink--active');

        const { container: containerFalse } = render(<WTag icon={AllInclusive} text="TagLink" isActive={false}/>);
        expect(containerFalse.firstChild).not.toHaveClass('tagLink--active');
    })
})
